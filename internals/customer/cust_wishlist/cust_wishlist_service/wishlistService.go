package custwishlistservice

import (
	"errors"

	custwishlist "github.com/ak-repo/ecommerce-gin/internals/customer/cust_wishlist"
	custwishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_wishlist/cust_wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type WishlistService struct {
	WishlistRepo custwishlistinterface.RepoInterface
}

func NewWishlistSevice(repo custwishlistinterface.RepoInterface) custwishlistinterface.ServiceInterface {
	return &WishlistService{WishlistRepo: repo}
}

// display all wishlist items
func (s *WishlistService) ListAllWishlistService(userID uint) (*custwishlist.WishlistDTO, error) {
	data, err := s.WishlistRepo.GetOrCreateWishlist(userID)
	if err != nil {
		return nil, err
	}

	var wishlistItem []custwishlist.WishlistItemDTO
	for _, w := range data.WishlistItems {
		item := custwishlist.WishlistItemDTO{
			ID:        w.ID,
			ProductID: w.ProductID,
			Product: struct {
				Name  string "json:\"name\""
				Price uint   "json:\"price\""
			}{
				Name:  w.Product.Title,
				Price: uint(w.Product.BasePrice),
			},
		}

		wishlistItem = append(wishlistItem, item)
	}

	wishlist := custwishlist.WishlistDTO{
		ID:            data.ID,
		UserID:        data.UserID,
		WishlistItems: wishlistItem,
	}

	return &wishlist, nil

}

// add to wishlist
func (s *WishlistService) AddToWishlistService(userID, productID uint) error {
	wishlist, err := s.WishlistRepo.GetOrCreateWishlist(userID)
	if err != nil {
		return err
	}

	// existing
	for _, items := range wishlist.WishlistItems {
		if items.ProductID == productID {
			return errors.New("product already in wishlist")
		}
	}

	wishlistItem := models.WishlistItem{
		WishlistID: wishlist.ID,
		ProductID:  productID,
	}
	return s.WishlistRepo.AddToWishlistItem(&wishlistItem)
}

// remove from wishlist
func (s *WishlistService) RemoveWishlistService(wishlistItemID uint) error {

	return s.WishlistRepo.DeleteWishlistItem(wishlistItemID)
}
