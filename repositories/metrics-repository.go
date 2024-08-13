package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
)

type MetricsRepository interface {
	GetTotalInvestedToDate() (float64, error)
	GetTotalCurrentlyInvested() (float64, error)
}

type metricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) MetricsRepository {
	return &metricsRepository{db}
}

func (r *metricsRepository) GetTotalInvestedToDate() (float64, error) {
	var total float64
	err := r.db.Model(&models.FixedTerm{}).Select("SUM(amount)").Scan(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *metricsRepository) GetTotalCurrentlyInvested() (float64, error) {
	var total float64
	err := r.db.Model(&models.FixedTerm{}).
		Select("SUM(amount)").
		Where("maturity_date > ?", time.Now()).
		Scan(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}
