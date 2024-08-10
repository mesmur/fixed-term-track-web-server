package repositories

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"gorm.io/gorm"
)

type InvestmentReturnRepository interface {
	FindByID(investmentId, returnId uint) (*models.InvestmentReturn, error)
	Create(investmentReturn *models.InvestmentReturn) (*models.InvestmentReturn, error)
}

type investmentReturnRepository struct {
	db *gorm.DB
}

func NewInvestmentReturnRepository(db *gorm.DB) InvestmentReturnRepository {
	return &investmentReturnRepository{db}
}

func (r *investmentReturnRepository) FindByID(investmentId, returnId uint) (*models.InvestmentReturn, error) {
	var investmentReturn models.InvestmentReturn
	err := r.db.Where("investment_id = ? AND id = ?", investmentId, returnId).First(&investmentReturn).Error

	if err != nil {
		return nil, err
	}

	return &investmentReturn, nil
}

func (r *investmentReturnRepository) Create(investmentReturn *models.InvestmentReturn) (*models.InvestmentReturn, error) {
	if err := r.db.Create(investmentReturn).Error; err != nil {
		return nil, err
	}
	return investmentReturn, nil
}
