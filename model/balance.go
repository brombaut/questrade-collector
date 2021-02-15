package model

import "fmt"

type Balance struct {
	Currency          string  `json:"currency"`
	Cash              float64 `json:"cash"`
	MarketValue       float64 `json:"marketValue"`
	TotalEquity       float64 `json:"totalEquity"`
	BuyingPower       float64 `json:"buyingPower"`
	MaintenanceExcess float64 `json:"maintenanceExcess"`
	IsRealTime        bool    `json:"isRealTime"`
}

type BalancesResponse struct {
	PerCurrencyBalances    []Balance `json:"perCurrencyBalances"`
	CombinedBalances       []Balance `json:"combinedBalances"`
	SodPerCurrencyBalances []Balance `json:"sodPerCurrencyBalances"`
	SodCombinedBalances    []Balance `json:"sodCombinedBalances"`
}

func (br BalancesResponse) TextOutput() string {
	result := fmt.Sprintf("%+v", br)
	return result
}
