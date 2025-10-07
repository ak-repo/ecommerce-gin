package custcheckoutservice

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_interface"
	custcheckout "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout"
	custcheckoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_interface"
	customerprofileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type CustomerCheckoutService struct {
	CartService    cartinterface.ServiceInterface
	AddressService customerprofileinterface.ServiceInterface
	UserRepo       authinterface.AuthRepoInterface
	CheckoutRepo   custcheckoutinterface.RepoInterface
}

func NewCustomerCheckoutService(customerCartService cartinterface.ServiceInterface, customerAddressService customerprofileinterface.ServiceInterface, userRepo authinterface.AuthRepoInterface, customerCheckoutRepo custcheckoutinterface.RepoInterface) custcheckoutinterface.ServiceInterface {
	return &CustomerCheckoutService{CartService: customerCartService, AddressService: customerAddressService, UserRepo: userRepo, CheckoutRepo: customerCheckoutRepo}
}

func (s *CustomerCheckoutService) CheckoutSummaryService(userID uint) (*custcheckout.CheckoutSummaryResponse, error) {

	if user, err := s.UserRepo.GetUserByID(userID); err != nil || user == nil {
		return nil, errors.New("can't find user,, checkout failed: " + err.Error())
	}

	var checkout custcheckout.CheckoutSummaryResponse
	// items
	cart, err := s.CartService.CustomerCartService(userID)
	if err != nil {
		return nil, err
	}
	checkout.Items = cart.Items

	// address
	address, err := s.AddressService.CustomerAddressService(userID)
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

func (s *CustomerCheckoutService) ProcessCheckoutService(req *custcheckout.CheckoutRequest) (*custcheckout.CheckoutResponse, error) {

	checkout, err := s.CheckoutSummaryService(req.UserID)
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
	if err := s.CartService.DeleteCartService(checkout.Items[0].CartID); err != nil {
		return nil, err
	}
	if err := s.CartService.DeleteCartitemBycartIDService(checkout.Items[0].CartID); err != nil {
		return nil, err
	}

	// res
	response := custcheckout.CheckoutResponse{
		OrderID:     order.ID,
		PaymentMode: req.PaymentMode,
		Amount:      order.TotalAmount,
		Status:      order.Status,
	}

	return &response, nil

}
