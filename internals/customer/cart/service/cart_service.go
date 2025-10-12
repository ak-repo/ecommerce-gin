package cartservice

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cart/cart_interface"
	"github.com/ak-repo/ecommerce-gin/internals/customer/cart/dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type service struct {
	CartRepo    cartinterface.Repository
	UserRepo    authinterface.Repository
	ProductRepo productinterface.Repository
}

func NewCartService(cartRepo cartinterface.Repository, userRepo authinterface.Repository, productRepo productinterface.Repository) cartinterface.Service {
	return &service{CartRepo: cartRepo, UserRepo: userRepo, ProductRepo: productRepo}
}

func (s *service) AddItem(userID uint, input *dto.AddToCartDTO) error {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	cart, err := s.CartRepo.GetOrCreateCart(user.ID)
	if err != nil {
		return err
	}

	for _, item := range cart.CartItems {
		if item.ProductID == input.ProductID {
			item.Quantity += input.Quantity
			item.Subtotal = item.Price * float64(item.Quantity)
			return s.CartRepo.UpdateCartItem(&item)
		}
	}

	product, err := s.ProductRepo.GetProductByID(input.ProductID)
	if err != nil {
		return err
	}

	newItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  input.Quantity,
		Price:     product.BasePrice,
		Subtotal:  product.BasePrice * float64(input.Quantity),
	}
	return s.CartRepo.CreateCartItem(&newItem)
}
func (s *service) GetUserCart(userID uint) (*dto.CartDTO, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	cart, err := s.CartRepo.GetOrCreateCart(user.ID)
	if err != nil {
		return nil, err
	}

	resp := &dto.CartDTO{
		CartID:   cart.ID,
		UserID:   user.ID,
		Currency: "INR",
	}

	for _, item := range cart.CartItems {
		resp.Items = append(resp.Items, dto.CartItemDTO{
			CartItemID:      item.ID,
			CartID:          item.CartID,
			ProductID:       item.Product.ID,
			ProductName:     item.Product.Title,
			ProductImageURL: item.Product.ImageURL,
			Price:           item.Price,
			Quantity:        item.Quantity,
			Subtotal:        item.Subtotal,
		})
		resp.Total += item.Subtotal
	}

	return resp, nil
}

func (s *service) UpdateQuantity(input *dto.UpdateCartItemDTO) error {
	item, err := s.CartRepo.GetCartItemByID(input.CartItemID)
	if err != nil || item == nil {
		return errors.New("cart item not found")
	}
	item.Quantity = input.Quantity
	item.Subtotal = item.Price * float64(input.Quantity)
	return s.CartRepo.UpdateCartItem(item)
}

func (s *service) RemoveItem(cartItemID uint) error {
	return s.CartRepo.DeleteCartItem(cartItemID)
}

func (s *service) DeleteCart(cartID uint) error {
	if err := s.CartRepo.DeleteCartItemsByCartID(cartID); err != nil {
		return err
	}
	return s.CartRepo.DeleteCart(cartID)
}
