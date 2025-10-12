package checkoutservice

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cart/cart_interface"
	checkoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/checkout_interface"
	checkoutdto "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/dto"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/profile/profile_interface"

	"github.com/ak-repo/ecommerce-gin/models"
)

type service struct {
	cartService    cartinterface.Service
	AddressService profileinterface.Service
	UserRepo       authinterface.Repository
	CheckoutRepo   checkoutinterface.Repository
}

func NewCheckoutService(cart cartinterface.Service, address profileinterface.Service, userRepo authinterface.Repository, checkout checkoutinterface.Repository) checkoutinterface.Service {
	return &service{cartService: cart, AddressService: address, UserRepo: userRepo, CheckoutRepo: checkout}
}

func (s *service) CheckoutSummary(userID uint) (*checkoutdto.CheckoutSummaryResponse, error) {

	if user, err := s.UserRepo.GetUserByID(userID); err != nil || user == nil {
		return nil, errors.New("can't find user,, checkout failed: " + err.Error())
	}

	var checkout checkoutdto.CheckoutSummaryResponse
	// items
	cart, err := s.cartService.GetUserCart(userID)
	if err != nil {
		return nil, err
	}
	checkout.Items = cart.Items

	// address
	address, err := s.AddressService.GetAddress(userID)
	if err != nil {
		return nil, err
	}
	checkout.Address = *address
	checkout.TotalItems = len(cart.Items)
	checkout.SubTotal = cart.Total
	checkout.ShippingFee = (checkout.SubTotal * 5) / 100
	checkout.GrandTotal = checkout.SubTotal + checkout.ShippingFee
	checkout.PaymentModes = []string{"COD", "ONLINE"}

	return &checkout, nil
}

func (s *service) ProcessCheckout(req *checkoutdto.CheckoutRequest) (*checkoutdto.CheckoutResponse, error) {

	checkout, err := s.CheckoutSummary(req.UserID)
	if err != nil {
		return nil, err
	}

	// order creation
	var order models.Order
	order.UserID = req.UserID
	order.Status = "pending"
	order.TotalAmount = checkout.GrandTotal
	order.ShippingAddressID = checkout.Address.ID

	if err := s.CheckoutRepo.OrderCreation(&order); err != nil {
		return nil, err
	}
	if order.ID == 0 {
		return nil, errors.New("order creation failed")

	}
	// order items
	var orderItems []models.OrderItem
	for _, i := range checkout.Items {
		var item models.OrderItem
		item.OrderID = order.ID
		item.ProductID = i.ProductID
		item.Quantity = i.Quantity
		item.UnitPrice = i.Price
		orderItems = append(orderItems, item)
	}

	if err := s.CheckoutRepo.OrderItemsCreation(orderItems); err != nil {
		return nil, err
	}
	// payment creation
	payment := models.Payment{
		OrderID:       order.ID,
		PaymentMethod: req.PaymentMode,
		Amount:        order.TotalAmount,
		Status:        "pending",
	}
	if err := s.CheckoutRepo.PaymentCreation(&payment); err != nil {
		return nil, err
	}

	// clear cart
	if err := s.cartService.DeleteCart(checkout.Items[0].CartID); err != nil {
		return nil, err
	}
	// res
	response := checkoutdto.CheckoutResponse{
		OrderID:     order.ID,
		PaymentMode: req.PaymentMode,
		Amount:      order.TotalAmount,
		Status:      order.Status,
	}

	return &response, nil

}
