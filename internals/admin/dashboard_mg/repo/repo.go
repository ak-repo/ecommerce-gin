package boardrepo

import (
	boardinter "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/dashboard_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewNewDashRepo(db *gorm.DB) boardinter.Repository {
	return &repository{db: db}
}

func (r *repository) TotalCount(model interface{}) (int64, error) {
	var count int64
	err := r.db.Model(&model).Count(&count).Error

	return count, err
}

func (r *repository) TotalRevenue() (float64, error) {
	var revenue float64
	err := r.db.Model(&models.Order{}).
		Where("status = ?", "completed").
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&revenue).Error

	return revenue, err
}

func (r *repository) Orders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}
