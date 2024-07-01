package domain

import "time"

type Ticker struct {
	Ltp       float64
	Symbol    string
	Type      string
	Timestamp time.Time
}
