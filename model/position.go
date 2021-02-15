package model

import "fmt"

type Position struct {
	Symbol             string  `json:"symbol"`
	SymbolId           float64 `json:"symbolId"`
	OpenQuantity       float64 `json:"openQuantity"`
	ClosedQuantity     float64 `json:"closedQuantity"`
	CurrentMarketValue float64 `json:"currentMarketValue"`
	CurrentPrice       float64 `json:"currentPrice"`
	AverageEntryPrice  float64 `json:"averageEntryPrice"`
	DayPnl             float64 `json:"dayPnl"`
	ClosedPnl          float64 `json:"closedPnl"`
	OpenPnl            float64 `json:"openPnl"`
	TotalCost          float64 `json:"totalCost"`
	IsRealTime         bool    `json:"isRealTime"`
	IsUnderReorg       bool    `json:"isUnderReorg"`
}

type PositionsResponse struct {
	Positions []Position `json:"positions"`
}

func (pr PositionsResponse) TextOutput() string {
	result := fmt.Sprintf("%+v", pr)
	return result
}
