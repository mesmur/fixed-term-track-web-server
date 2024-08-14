package models

import "time"

type FixedTermReturn struct {
	Base
	FixedTermID      uint      `json:"fixed_term_id"` // Foreign key (belongs to), tag `index` will create index for this column
	Interest         float64   `json:"interest"`
	WithholdingTax   float64   `json:"withholding_tax"`
	Amount           float64   `json:"amount"`            // Interest - WithholdingTax
	Currency         string    `json:"currency"`          // Currency of the amount
	AmountPercentage float64   `json:"amount_percentage"` // Amount / Principal Amount (from FixedTerm)
	Date             time.Time `json:"date"`              // Payout date
}
