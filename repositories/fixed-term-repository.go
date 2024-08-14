package repositories

import (
	"gorm.io/gorm"

	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
)

type FixedTermRepository interface {
	FindByID(id uint) (*models.FixedTerm, error)
	Create(fixedTerm *models.FixedTerm) error
	Update(fixedTerm *models.FixedTerm) error
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

func (r *fixedTermRepository) Create(fixedTerm *models.FixedTerm) error {
	if err := r.db.Create(fixedTerm).Error; err != nil {
		return err
	}
	return nil
}

func (r *fixedTermRepository) Update(fixedTerm *models.FixedTerm) error {
	if err := r.db.Save(fixedTerm).Error; err != nil {
		return err
	}
	return nil
}
