package orderservice

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	orderrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/orderRepository"
)

type OrderService interface {
	GetAllOrdersService() ([]models.Order, error)
	GetOrderByIDService(id string) (*models.Order, error)
}

type orderService struct {
	orderRepo orderrepository.OrderRepo
	cfg       *config.Config
}

func NewOrderService(orderRepo orderrepository.OrderRepo, cfg *config.Config) OrderService {
	return &orderService{orderRepo: orderRepo, cfg: cfg}
}

// All orders fetch
func (s *orderService) GetAllOrdersService() ([]models.Order, error) {

	orders, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil

}

// GEt one order
func (s *orderService) GetOrderByIDService(id string) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
