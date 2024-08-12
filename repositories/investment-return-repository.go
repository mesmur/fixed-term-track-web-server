package repositories

import (
	"gorm.io/gorm"

	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
)

type FixedTermReturnRepository interface {
	FindByID(fixedTermID, returnId uint) (*models.FixedTermReturn, error)
	Create(fixedTermReturn *models.FixedTermReturn) (*models.FixedTermReturn, error)
}

type fixedTermReturnRepository struct {
	db *gorm.DB
}

func NewFixedTermReturnRepository(db *gorm.DB) FixedTermReturnRepository {
	return &fixedTermReturnRepository{db}
}

func (r *fixedTermReturnRepository) FindByID(fixedTermId, returnId uint) (*models.FixedTermReturn, error) {
	var fixedTermReturn models.FixedTermReturn
	err := r.db.Where("fixed_term_id = ? AND id = ?", fixedTermId, returnId).First(&fixedTermReturn).Error

	if err != nil {
		return nil, err
	}

	return &fixedTermReturn, nil
}

func (r *fixedTermReturnRepository) Create(fixedTermReturn *models.FixedTermReturn) (*models.FixedTermReturn, error) {
	if err := r.db.Create(fixedTermReturn).Error; err != nil {
		return nil, err
	}
	return fixedTermReturn, nil
}
