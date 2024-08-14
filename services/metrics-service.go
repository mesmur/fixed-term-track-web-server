package services

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
)

type MetricsService interface {
	GetTotalInvestedToDate() (*models.Metric, error)
	GetTotalCurrentlyInvested() (*models.Metric, error)
	GetTotalMaturingInMonths(months int) (*models.Metric, error)
	GetTotalReturnsToDate() (*models.Metric, error)
	GetTotalReturnsThisYear() (*models.Metric, error)
}

type metricsService struct {
	repo repositories.MetricsRepository
}

func NewMetricsService(repo repositories.MetricsRepository) MetricsService {
	return &metricsService{repo}
}

func (s *metricsService) GetTotalInvestedToDate() (*models.Metric, error) {
	return s.repo.GetTotalInvestedToDate()
}

func (s *metricsService) GetTotalCurrentlyInvested() (*models.Metric, error) {
	return s.repo.GetTotalCurrentlyInvested()
}

func (s *metricsService) GetTotalMaturingInMonths(months int) (*models.Metric, error) {
	return s.repo.GetTotalMaturingInMonths(months)
}

func (s *metricsService) GetTotalReturnsToDate() (*models.Metric, error) {
	return s.repo.GetTotalReturnsToDate()
}

func (s *metricsService) GetTotalReturnsThisYear() (*models.Metric, error) {
	return s.repo.GetTotalReturnsThisYear()
}
