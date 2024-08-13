package services

import (
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
)

type MetricsService interface {
	GetTotalInvestedToDate() (float64, error)
	GetTotalCurrentlyInvested() (float64, error)
}

type metricsService struct {
	repo repositories.MetricsRepository
}

func NewMetricsService(repo repositories.MetricsRepository) MetricsService {
	return &metricsService{repo}
}

func (s *metricsService) GetTotalInvestedToDate() (float64, error) {
	return s.repo.GetTotalInvestedToDate()
}

func (s *metricsService) GetTotalCurrentlyInvested() (float64, error) {
	return s.repo.GetTotalCurrentlyInvested()
}
