package repositories

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"gorm.io/gorm"
)

type FixedTermRepository interface {
	FindByID(id uint) (*models.FixedTerm, error)
	Create(fixedTerm *models.FixedTerm) (*models.FixedTerm, error)
	Update(fixedTerm *models.FixedTerm) (*models.FixedTerm, error)
}

type fixedTermRepository struct {
	db *gorm.DB
}

func NewFixedTermRepository(db *gorm.DB) FixedTermRepository {
	return &fixedTermRepository{db}
}

func (r *fixedTermRepository) FindByID(id uint) (*models.FixedTerm, error) {
	var fixedTerm models.FixedTerm
	if err := r.db.Preload("FixedTermReturns").First(&fixedTerm, id).Error; err != nil {
		return nil, err
	}
	return &fixedTerm, nil
}

func (r *fixedTermRepository) Create(fixedTerm *models.FixedTerm) (*models.FixedTerm, error) {
	if err := r.db.Create(fixedTerm).Error; err != nil {
		return nil, err
	}
	return fixedTerm, nil
}

func (r *fixedTermRepository) Update(fixedTerm *models.FixedTerm) (*models.FixedTerm, error) {
	if err := r.db.Save(fixedTerm).Error; err != nil {
		return nil, err
	}
	return fixedTerm, nil
}
