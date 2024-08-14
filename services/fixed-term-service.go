package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
)

type FixedTermService interface {
	FindByID(id uint) (*models.FixedTerm, error)
	Create(fixedTerm *models.FixedTerm) error
	Update(fixedTerm *models.FixedTerm) error
	FindReturnByID(fixedTermID, returnId uint) (*models.FixedTermReturn, error)
	CreateReturn(fixedTermReturn *models.FixedTermReturn) error
}

type fixedTermService struct {
	fixedTermRepository       repositories.FixedTermRepository
	fixedTermReturnRepository repositories.FixedTermReturnRepository
	eventRepository           repositories.EventRepository
}

func NewFixedTermService(fixedTermRepository repositories.FixedTermRepository, fixedTermReturnRepository repositories.FixedTermReturnRepository, eventRepository repositories.EventRepository) FixedTermService {
	return &fixedTermService{fixedTermRepository, fixedTermReturnRepository, eventRepository}
}

func (s *fixedTermService) FindByID(id uint) (*models.FixedTerm, error) {
	return s.fixedTermRepository.FindByID(id)
}

func (s *fixedTermService) Create(fixedTerm *models.FixedTerm) error {
	fixedTerm.MaturityDate = calculateMaturityDate(fixedTerm)

	err := s.fixedTermRepository.Create(fixedTerm)

	if err != nil {
		return err
	}

	if fixedTerm.ReturnType == "MATURITY" {
		err = s.scheduleMaturityEvent(fixedTerm)
	} else if fixedTerm.ReturnType == "MONTHLY" {
		err = s.scheduleMonthlyEvents(fixedTerm)
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *fixedTermService) Update(fixedTerm *models.FixedTerm) error {
	return s.fixedTermRepository.Update(fixedTerm)
}

func (s *fixedTermService) FindReturnByID(fixedTermID, returnId uint) (*models.FixedTermReturn, error) {
	return s.fixedTermReturnRepository.FindByID(fixedTermID, returnId)
}

func (s *fixedTermService) CreateReturn(fixedTermReturn *models.FixedTermReturn) error {
	var fixedTerm, err = s.fixedTermRepository.FindByID(fixedTermReturn.FixedTermID)

	if err != nil {
		return err
	}

	if fixedTermReturn.Date.Before(fixedTerm.OpenDate) {
		return errors.New("A return cannot be created before the open date of the Fixed Term")
	}

	if fixedTermReturn.Date.After(fixedTerm.MaturityDate) {
		return errors.New("A return cannot be created after the maturity date of the Fixed Term")
	}

	fixedTermReturn.Amount = fixedTermReturn.Interest - fixedTermReturn.WithholdingTax
	fixedTermReturn.AmountPercentage = (fixedTermReturn.Amount / fixedTerm.Amount) * 100
	fixedTermReturn.Currency = fixedTerm.Currency

	return s.fixedTermReturnRepository.Create(fixedTermReturn)
}

// create a function to do this
func calculateMaturityDate(fixedTerm *models.FixedTerm) time.Time {
	// Calculate the first day of the next month using the open date
	firstOfNextMonth := time.Date(fixedTerm.OpenDate.Year(), fixedTerm.OpenDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	// Add the period months to the first day of the next month
	maturityDate := firstOfNextMonth.AddDate(0, int(fixedTerm.Period), 0)

	return maturityDate
}

func (s *fixedTermService) scheduleMonthlyEvents(fixedTerm *models.FixedTerm) error {
	// The fixedTerm would be 'active' in the first of the next month
	// So the first return date would be one month after that
	firstReturnDate := time.Date(fixedTerm.OpenDate.Year(), fixedTerm.OpenDate.Month()+2, 1, 0, 0, 0, 0, time.UTC)

	// Schedule events for each month until the maturity date
	for date := firstReturnDate; !date.After(fixedTerm.MaturityDate); date = date.AddDate(0, 1, 0) {
		// Assumes monthly returns are constant throughout the period - unlikely in real life
		estimatedReturn := (fixedTerm.Amount * (fixedTerm.APY / 100)) / 12
		msg := fmt.Sprintf("Check on your Term FixedTerm %s from %s.\nThere should be a monthly return of %.2f %s (estimated) :)\n\nFixed Term ID: %s", fixedTerm.BankTermID, fixedTerm.Bank, estimatedReturn, fixedTerm.Currency, strconv.Itoa(int(fixedTerm.ID)))

		event := models.Event{
			ResourceID:    fixedTerm.ID,
			ScheduledTime: date,
			EventType:     "MONTHLY_RETURN_NOTIFICATION",
			Status:        "SCHEDULED",
			Message:       msg,
		}

		err := s.eventRepository.Create(&event)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *fixedTermService) scheduleMaturityEvent(fixedTerm *models.FixedTerm) error {
	estimatedReturn := (fixedTerm.Amount * (fixedTerm.APY / 100)) * float64(fixedTerm.Period/12)
	msg := fmt.Sprintf("Check on your Term FixedTerm %s from %s.\nThere should be a maturity return of %.2f %s (estimated) :)\n\nFixed Term ID: %s", fixedTerm.BankTermID, fixedTerm.Bank, estimatedReturn, fixedTerm.Currency, strconv.Itoa(int(fixedTerm.ID)))
	event := models.Event{
		ResourceID:    fixedTerm.ID,
		ScheduledTime: fixedTerm.MaturityDate,
		EventType:     "MATURITY_RETURN_NOTIFICATION",
		Status:        "SCHEDULED",
		Message:       msg,
	}

	err := s.eventRepository.Create(&event)
	if err != nil {
		return err
	}

	return nil
}
