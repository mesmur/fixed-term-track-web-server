package models

import "time"

type Investment struct {
	Base
	Bank              string    `json:"bank"`
	DepositID         string    `json:"deposit_id"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	TermPeriod        uint      `json:"term_period"`  // Number of months
	DepositType       string    `json:"deposit_type"` // 'MONTHLY' or 'MATURITY'
	APY               float64   `json:"apy"`          // APY = Annual Percentage Yield
	OpenDate          time.Time `json:"open_date" gorm:"index"`
	MaturityDate      time.Time `json:"maturity_date"`
	AutomaticRenewal  bool      `json:"automatic_renewal"`
	IsComplete        bool      `json:"is_complete"`
	InvestmentReturns []InvestmentReturn
}
