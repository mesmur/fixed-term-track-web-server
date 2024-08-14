package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
)

type MetricsRepository interface {
	GetTotalInvestedToDate() (*models.Metric, error)
	GetTotalCurrentlyInvested() (*models.Metric, error)
	GetTotalMaturingInMonths(months int) (*models.Metric, error)
	GetTotalReturnsToDate() (*models.Metric, error)
	GetTotalReturnsThisYear() (*models.Metric, error)
}

type metricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) MetricsRepository {
	return &metricsRepository{db}
}

func (r *metricsRepository) GetTotalInvestedToDate() (*models.Metric, error) {
	var metric models.Metric
	var data []models.MetricData

	err := r.db.Model(&models.FixedTerm{}).
		Select("SUM(amount) as Amount, currency").
		Group("currency").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	logger.Sugar.Infof("Data: %+v", data)

	metric.Data = data

	return &metric, nil
}

func (r *metricsRepository) GetTotalCurrentlyInvested() (*models.Metric, error) {
	var metric models.Metric
	var data []models.MetricData

	err := r.db.Model(&models.FixedTerm{}).
		Select("SUM(amount) as Amount, currency").
		Where("maturity_date > ?", time.Now()).
		Group("currency").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	metric.Data = data

	return &metric, nil
}

func (r *metricsRepository) GetTotalMaturingInMonths(months int) (*models.Metric, error) {
	var metric models.Metric
	var data []models.MetricData
	start := time.Now()
	end := time.Now().AddDate(0, months, 0)

	err := r.db.Model(&models.FixedTerm{}).
		Select("SUM(amount) as Amount, currency").
		Where("maturity_date BETWEEN ? AND ?", start, end).
		Group("currency").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	metric.Data = data
	metric.FromDate = &start
	metric.ToDate = &end

	return &metric, nil
}

func (r *metricsRepository) GetTotalReturnsToDate() (*models.Metric, error) {
	var metric models.Metric
	var data []models.MetricData

	err := r.db.Model(&models.FixedTermReturn{}).
		Select("SUM(amount) as Amount, currency").
		Group("currency").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	metric.Data = data

	return &metric, nil
}

func (r *metricsRepository) GetTotalReturnsThisYear() (*models.Metric, error) {
	var metric models.Metric
	var data []models.MetricData

	start := getStartOfYear()
	end := time.Now()

	err := r.db.Model(&models.FixedTermReturn{}).
		Select("SUM(amount) as Amount, currency").
		Where("date BETWEEN ? AND ?", start, end).
		Group("currency").
		Scan(&data).Error

	if err != nil {
		return nil, err
	}

	metric.Data = data
	metric.FromDate = &start
	metric.ToDate = &end

	return &metric, nil
}

func getStartOfYear() time.Time {
	now := time.Now()
	startOfYear := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	return startOfYear
}
