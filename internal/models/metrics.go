package models

import (
	"time"
)

type Metric struct {
	Data     []MetricData `json:"data"`
	FromDate *time.Time   `json:"from_date"`
	ToDate   *time.Time   `json:"to_date"`
}

type MetricData struct {
	Amount   float64 `json:"amount" gorm:"column:amount"`
	Currency string  `json:"currency" gorm:"column:currency"`
}

func (m *Metric) AddData(amount float64, currency string) {
	m.Data = append(m.Data, MetricData{Amount: amount, Currency: currency})
}
