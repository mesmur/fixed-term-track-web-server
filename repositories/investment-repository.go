package repositories

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"gorm.io/gorm"
)

type InvestmentRepository interface {
	FindByID(id uint) (*models.Investment, error)
	Create(investment *models.Investment) (*models.Investment, error)
	Update(investment *models.Investment) (*models.Investment, error)
}

type investmentRepository struct {
	db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) InvestmentRepository {
	return &investmentRepository{db}
}

func (r *investmentRepository) FindByID(id uint) (*models.Investment, error) {
	var investment models.Investment
	if err := r.db.Preload("InvestmentReturns").First(&investment, id).Error; err != nil {
		return nil, err
	}
	return &investment, nil
}

func (r *investmentRepository) Create(investment *models.Investment) (*models.Investment, error) {
	if err := r.db.Create(investment).Error; err != nil {
		return nil, err
	}
	return investment, nil
}

func (r *investmentRepository) Update(investment *models.Investment) (*models.Investment, error) {
	if err := r.db.Save(investment).Error; err != nil {
		return nil, err
	}
	return investment, nil
}
