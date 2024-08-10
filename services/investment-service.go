package services

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
	"time"
)

type InvestmentService interface {
	FindByID(id uint) (*models.Investment, error)
	Create(investment *models.Investment) (*models.Investment, error)
	Update(investment *models.Investment) (*models.Investment, error)
	FindReturnByID(investmentId, returnId uint) (*models.InvestmentReturn, error)
	CreateInvestmentReturn(investmentReturn *models.InvestmentReturn) (*models.InvestmentReturn, error)
}

type investmentService struct {
	investmentRepo       repositories.InvestmentRepository
	investmentReturnRepo repositories.InvestmentReturnRepository
	eventRepo            repositories.EventRepository
}

func NewInvestmentService(investmentRepo repositories.InvestmentRepository, investmentReturnRepo repositories.InvestmentReturnRepository, eventRepository repositories.EventRepository) InvestmentService {
	return &investmentService{investmentRepo, investmentReturnRepo, eventRepository}
}

func (s *investmentService) FindByID(id uint) (*models.Investment, error) {
	return s.investmentRepo.FindByID(id)
}

func (s *investmentService) Create(investment *models.Investment) (*models.Investment, error) {
	investment.MaturityDate = calculateMaturityDate(investment)
	logger.Sugar.Info("Maturity date: ", investment.MaturityDate)

	investment, err := s.investmentRepo.Create(investment)

	if err != nil {
		return nil, err
	}

	if investment.DepositType == "MATURITY" {
		err = s.scheduleMaturityEvent(investment)
	} else if investment.DepositType == "MONTHLY" {
		err = s.scheduleMonthlyEvents(investment)
	}

	if err != nil {
		return nil, err
	}

	return investment, nil
}

func (s *investmentService) Update(investment *models.Investment) (*models.Investment, error) {
	return s.investmentRepo.Update(investment)
}

func (s *investmentService) FindReturnByID(investmentId, returnId uint) (*models.InvestmentReturn, error) {
	return s.investmentReturnRepo.FindByID(investmentId, returnId)
}

func (s *investmentService) CreateInvestmentReturn(investmentReturn *models.InvestmentReturn) (*models.InvestmentReturn, error) {
	// Manually calculate teh amount and return
	var investment, err = s.investmentRepo.FindByID(investmentReturn.InvestmentID)

	if err != nil {
		return nil, err
	}

	investmentReturn.Amount = investmentReturn.Interest - investmentReturn.WithholdingTax
	investmentReturn.AmountPercentage = investmentReturn.Amount / investment.Amount

	return s.investmentReturnRepo.Create(investmentReturn)
}

// create a function to do this
func calculateMaturityDate(investment *models.Investment) time.Time {
	// Calculate the first day of the next month using the open date
	firstOfNextMonth := time.Date(investment.OpenDate.Year(), investment.OpenDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	// Add the period months to the first day of the next month
	maturityDate := firstOfNextMonth.AddDate(0, int(investment.TermPeriod), 0)

	return maturityDate
}

func (s *investmentService) scheduleMonthlyEvents(investment *models.Investment) error {

	// The investment would be 'active' in the first of the next month
	// So the first return date would be one month after that
	firstReturnDate := time.Date(investment.OpenDate.Year(), investment.OpenDate.Month()+2, 1, 0, 0, 0, 0, time.UTC)

	// Schedule events for each month until the maturity date and one more
	for date := firstReturnDate; !date.After(investment.MaturityDate); date = date.AddDate(0, 1, 0) {
		event := models.Event{
			ResourceID:    investment.ID,
			ScheduledTime: date,
			EventType:     "MONTHLY_RETURN_NOTIFICATION",
			Status:        "SCHEDULED",
		}

		err := s.eventRepo.Create(&event)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *investmentService) scheduleMaturityEvent(investment *models.Investment) error {
	event := models.Event{
		ResourceID:    investment.ID,
		ScheduledTime: investment.MaturityDate,
		EventType:     "MATURITY_RETURN_NOTIFICATION",
		Status:        "SCHEDULED",
	}

	err := s.eventRepo.Create(&event)

	if err != nil {
		return err
	}

	return nil
}
