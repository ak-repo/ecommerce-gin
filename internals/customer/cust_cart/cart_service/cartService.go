package cartservice

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	custcart "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart"
	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_interface"
	custproductinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type CartService struct {
	CartRepo    cartinterface.RepoInterface
	UserRepo    authinterface.AuthRepoInterface
	ProductRepo custproductinterface.RepoInterface
}

func NewCartRepository(cartRepo cartinterface.RepoInterface, userRepo authinterface.AuthRepoInterface, productRepo custproductinterface.RepoInterface) cartinterface.ServiceInterface {
	return &CartService{CartRepo: cartRepo, UserRepo: userRepo, ProductRepo: productRepo}
}

// add item into user cart
func (s *CartService) AddtoCartService(userID uint, addtoCart *custcart.AddToCartDTO) error {

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	cart, err := s.CartRepo.GetorCreateCart(user.ID)

	if err == nil && cart.CartItems != nil {
		for _, cartItem := range cart.CartItems {
			if cartItem.ProductID == addtoCart.ProductID {
				cartItem.Quantity += addtoCart.Quantity
				cartItem.Subtotal = cartItem.Price * float64(cartItem.Quantity)
				return s.CartRepo.UpdateCartItem(&cartItem)
			}
		}

	}

	// add new cart item
	product, err := s.ProductRepo.GetProductByID(addtoCart.ProductID)
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

	return s.CartRepo.CreateCartItem(&newItem)
}

// display all cart items
func (s *CartService) CustomerCartService(userID uint) (*custcart.CartDTO, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err

	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	cart, err := s.CartRepo.GetorCreateCart(user.ID)

	if err != nil {
		return nil, err
	}

	fullCart := custcart.CartDTO{
		CartID: cart.ID,
		UserID: user.ID,
	}

	// items chnage
	for _, item := range cart.CartItems {
		product, err := s.ProductRepo.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		cartItem := custcart.CartItemDTO{
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
func (s *CartService) UpdateQuantityService(updatedCart *custcart.UpdateCartItemDTO) error {

	cartItem, err := s.CartRepo.GetCartItemByID(updatedCart.CartItemID)
	if err != nil {
		return err
	}
	cartItem.Quantity = updatedCart.Quantity
	cartItem.Subtotal = cartItem.Price * float64(updatedCart.Quantity)
	return s.CartRepo.UpdateCartItem(cartItem)
}

// remove cart item from cart
func (s *CartService) RemoveCartItemService(cartItemID uint) error {

	return s.CartRepo.DeleteCartItem(cartItemID)
}

// delete cart
func (s *CartService) DeleteCartService(cartID uint) error {

	return s.CartRepo.DeleteCart(cartID)
}

// delete cartitme by cart id
func (s *CartService) DeleteCartitemBycartIDService(cartID uint) error {

	return s.CartRepo.DeleteCartItemBycartID(cartID)
}