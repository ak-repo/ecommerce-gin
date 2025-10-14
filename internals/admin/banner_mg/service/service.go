package bannerservice

import (
	bannerinter "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/banner_interface"
	bannerdto "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
)

type service struct {
	BannerRepo bannerinter.Repository
}

func NewBannerServiceMG(repo bannerinter.Repository) bannerinter.Service {

	return &service{BannerRepo: repo}
}

func (s *service) Create(req *bannerdto.CreateBannerRequest) (uint, error) {

	banner := models.Banner{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		IsActive:    req.IsActive,
	}

	if err := s.BannerRepo.Create(&banner); err != nil {
		return 0, err
	}

	return banner.ID, nil
}

func (s *service) Update(req *bannerdto.UpdateBannerRequest) error {
	existingBanner, err := s.BannerRepo.GetBannerByID(req.ID)
	if err != nil {
		return err
	}

	if req.Title != "" {
		existingBanner.Title = req.Title
	}
	if req.Description != "" {
		existingBanner.Description = req.Description
	}
	if req.ImageURL != "" {
		existingBanner.ImageURL = req.ImageURL
	}

	existingBanner.IsActive = req.IsActive

	return s.BannerRepo.Update(existingBanner)
}

func (s *service) Delete(bannerID uint) error {
	return s.BannerRepo.Delete(bannerID)
}

func (s *service) GetAllBanners() ([]bannerdto.BannerResponse, error) {
	data, err := s.BannerRepo.GetAllBanners()
	if err != nil {
		return nil, err
	}
	var banners []bannerdto.BannerResponse
	for _, i := range data {

		banner := bannerdto.BannerResponse{
			ID:          i.ID,
			Title:       i.Title,
			Description: i.Description,
			ImageURL:    i.ImageURL,
			IsActive:    i.IsActive,
			CreatedAt:   i.CreatedAt,
			UpdatedAt:   i.UpdatedAt,
		}
		banners = append(banners, banner)
	}
	return banners, err
}

func (s *service) GetBannerByID(bannerID uint) (*bannerdto.BannerResponse, error) {

	data, err := s.BannerRepo.GetBannerByID(bannerID)
	if err != nil {
		return nil, err
	}

	banner := bannerdto.BannerResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		ImageURL:    data.ImageURL,
		IsActive:    data.IsActive,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return &banner, nil

}
