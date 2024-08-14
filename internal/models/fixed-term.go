package models

import "time"

type FixedTerm struct {
	Base
	Bank             string    `json:"bank" gorm:"uniqueIndex:constraint_bank_x_bank_term_id"`
	BankTermID       string    `json:"bank_term_id" gorm:"uniqueIndex:constraint_bank_x_bank_term_id"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	Period           uint      `json:"period"`      // Number of months
	Type             string    `json:"type"`        // "FIXED_DEPOSIT", "SHARE_CERTIFICATE"
	ReturnType       string    `json:"return_type"` // 'MONTHLY' or 'MATURITY'
	APY              float64   `json:"apy"`         // APY = Annual Percentage Yield
	OpenDate         time.Time `json:"open_date" gorm:"index"`
	MaturityDate     time.Time `json:"maturity_date"`
	FixedTermReturns []FixedTermReturn
}
