package pagessvc

import (
	pagesdto "github.com/ak-repo/ecommerce-gin/internals/customer/pages/dto"
	pagesinter "github.com/ak-repo/ecommerce-gin/internals/customer/pages/pages_interface"
)

type service struct {
	PagesRepo pagesinter.Repository
}

func NewPagesService(repo pagesinter.Repository) pagesinter.Service {
	return &service{PagesRepo: repo}
}

func (s *service) GetBanners() ([]pagesdto.BannerDTO, error) {

	data, err := s.PagesRepo.GetBanners()
	if err != nil {
		return nil, err
	}

	var banners []pagesdto.BannerDTO
	for _, i := range data {
		banner := pagesdto.BannerDTO{
			ID:          i.ID,
			Title:       i.Title,
			Description: i.Description,
			ImageURL:    i.ImageURL,
		}
		banners = append(banners, banner)
	}
	return banners, nil
}
