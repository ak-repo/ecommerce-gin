package cartservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	cartrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/cartRepository"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
)

type CartService interface {
	AddtoCartService(userID uint, addtoCart *dto.AddToCartDTO) error
	UserCartService(userID uint) (*dto.CartDTO, error)
	UpdateQuantityService(updatedCart *dto.UpdateCartItemDTO) error
	RemoveCartItemService(cartItemID uint) error
}

type cartService struct {
	cartRepository    cartrepository.CartRepository
	userRepository    userrepository.UserRepo
	productRepository productrepository.ProductRepo
}

func NewCartRepository(userRepo userrepository.UserRepo, cartRepo cartrepository.CartRepository, productRepo productrepository.ProductRepo) CartService {
	return &cartService{userRepository: userRepo, cartRepository: cartRepo, productRepository: productRepo}
}

// add item into user cart
func (s *cartService) AddtoCartService(userID uint, addtoCart *dto.AddToCartDTO) error {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	cart, err := s.cartRepository.GetorCreateCart(user.ID)

	if err == nil && cart.CartItems != nil {
		for _, cartItem := range cart.CartItems {
			if cartItem.ProductID == addtoCart.ProductID {
				cartItem.Quantity += addtoCart.Quantity
				cartItem.Subtotal = cartItem.Price * float64(cartItem.Quantity)
				return s.cartRepository.UpdateCartItem(&cartItem)
			}
		}

	}

	// add new cart item
	product, err := s.productRepository.GetProductByID(addtoCart.ProductID)
	if err != nil {
		return err
	}
	newItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: product.ID,
		Quantity:  addtoCart.Quantity,
		Price:     product.BasePrice,
		Subtotal:  float64(addtoCart.Quantity) * product.BasePrice,
	}

	return s.cartRepository.CreateCartItem(&newItem)
}

// display all cart items
func (s *cartService) UserCartService(userID uint) (*dto.CartDTO, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err

	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	cart, err := s.cartRepository.GetorCreateCart(user.ID)

	if err != nil {
		return nil, err
	}

	fullCart := dto.CartDTO{
		CartID: cart.ID,
		UserID: user.ID,
	}

	// items chnage
	for _, item := range cart.CartItems {
		product, err := s.productRepository.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		cartItem := dto.CartItemDTO{
			CartItemID:      item.ID,
			CartID:          item.CartID,
			ProductID:       item.ProductID,
			Product:         product.Title,
			ProductImageURL: product.ImageURL,
			Quantity:        item.Quantity,
			Subtotal:        item.Subtotal,
			Price:           item.Price,
		}
		fullCart.Items = append(fullCart.Items, cartItem)
		fullCart.Total += item.Subtotal // total amount
	}

	return &fullCart, nil

}

// update cart quantity
func (s *cartService) UpdateQuantityService(updatedCart *dto.UpdateCartItemDTO) error {

	cartItem, err := s.cartRepository.GetCartItemByID(updatedCart.CartItemID)
	if err != nil {
		return err
	}
	cartItem.Quantity = updatedCart.Quantity
	cartItem.Subtotal = cartItem.Price * float64(updatedCart.Quantity)
	return s.cartRepository.UpdateCartItem(cartItem)
}

// remove cart item from cart
func (s *cartService) RemoveCartItemService(cartItemID uint) error {

	return s.cartRepository.DeleteCartItem(cartItemID)
}
