package models

import "time"

type InvestmentReturn struct {
	Base
	InvestmentID     uint      `json:"investment_id"` // Foreign key (belongs to), tag `index` will create index for this column
	Interest         float64   `json:"interest"`
	WithholdingTax   float64   `json:"withholding_tax"`
	Amount           float64   `json:"amount"`            // Interest - WithholdingTax
	AmountPercentage float64   `json:"amount_percentage"` // Amount / Principal Amount (from Investment)
	Date             time.Time `json:"date"`
}
