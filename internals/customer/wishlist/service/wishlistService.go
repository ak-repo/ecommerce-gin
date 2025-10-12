package wishlistservice

import (
	"errors"

	wishlistdto "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/wishlist_dto"
	wishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type service struct {
	WishlistRepo wishlistinterface.Repository
}

func NewWishlistSevice(repo wishlistinterface.Repository) wishlistinterface.Service {
	return &service{WishlistRepo: repo}
}

// display all wishlist items
func (s *service) List(userID uint) (*wishlistdto.WishlistDTO, error) {
	data, err := s.WishlistRepo.List(userID)
	if err != nil {
		return nil, err
	}

	var wishlistItem []wishlistdto.WishlistItemDTO
	for _, w := range data.WishlistItems {
		item := wishlistdto.WishlistItemDTO{
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

	wishlist := wishlistdto.WishlistDTO{
		ID:            data.ID,
		UserID:        data.UserID,
		WishlistItems: wishlistItem,
	}

	return &wishlist, nil

}


func (s *service) Add(userID, productID uint) error {
	wishlist, err := s.WishlistRepo.List(userID)
	if err != nil {
		return err
	}

	for _, items := range wishlist.WishlistItems {
		if items.ProductID == productID {
			return errors.New("product already in wishlist")
		}
	}

	wishlistItem := models.WishlistItem{
		WishlistID: wishlist.ID,
		ProductID:  productID,
	}
	return s.WishlistRepo.Add(&wishlistItem) 
}


func (s *service) Remove(wishlistItemID uint) error {
	return s.WishlistRepo.Remove(wishlistItemID)
}
