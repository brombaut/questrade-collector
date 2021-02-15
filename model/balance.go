package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

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

type CSVBalances struct {
	Rows []CSVBalance
}

type CSVBalance struct {
	Date              time.Time
	AccountType       string
	AccountNumber     string
	BalanceType       string
	Currency          string
	Cash              float64
	MarketValue       float64
	TotalEquity       float64
	BuyingPower       float64
	MaintenanceExcess float64
	IsRealTime        bool
}

func CSVBalanceHeaders() []string {
	return []string{
		"Date",
		"AccountType",
		"AccountNumber",
		"BalanceType",
		"Currency",
		"Cash",
		"MarketValue",
		"TotalEquity",
		"BuyingPower",
		"MaintenanceExcess",
		"IsRealTime",
	}
}

func (csvB *CSVBalances) FromBalances(balances BalancesResponse, account Account) {
	for _, b := range balances.PerCurrencyBalances {
		var row CSVBalance
		row.Date = time.Now()
		row.AccountType = account.Type
		row.AccountNumber = account.Number
		row.BalanceType = "PerCurrencyBalances"
		row.Currency = b.Currency
		row.Cash = b.Cash
		row.MarketValue = b.MarketValue
		row.TotalEquity = b.TotalEquity
		row.BuyingPower = b.BuyingPower
		row.MaintenanceExcess = b.MaintenanceExcess
		row.IsRealTime = b.IsRealTime
		csvB.Rows = append(csvB.Rows, row)
	}

	for _, b := range balances.CombinedBalances {
		var row CSVBalance
		row.Date = time.Now()
		row.AccountType = account.Type
		row.AccountNumber = account.Number
		row.BalanceType = "CombinedBalances"
		row.Currency = b.Currency
		row.Cash = b.Cash
		row.MarketValue = b.MarketValue
		row.TotalEquity = b.TotalEquity
		row.BuyingPower = b.BuyingPower
		row.MaintenanceExcess = b.MaintenanceExcess
		row.IsRealTime = b.IsRealTime
		csvB.Rows = append(csvB.Rows, row)
	}

	for _, b := range balances.SodPerCurrencyBalances {
		var row CSVBalance
		row.Date = time.Now()
		row.AccountType = account.Type
		row.AccountNumber = account.Number
		row.BalanceType = "SodPerCurrencyBalances"
		row.Currency = b.Currency
		row.Cash = b.Cash
		row.MarketValue = b.MarketValue
		row.TotalEquity = b.TotalEquity
		row.BuyingPower = b.BuyingPower
		row.MaintenanceExcess = b.MaintenanceExcess
		row.IsRealTime = b.IsRealTime
		csvB.Rows = append(csvB.Rows, row)
	}

	for _, b := range balances.SodCombinedBalances {
		var row CSVBalance
		row.Date = time.Now()
		row.AccountType = account.Type
		row.AccountNumber = account.Number
		row.BalanceType = "SodCombinedBalances"
		row.Currency = b.Currency
		row.Cash = b.Cash
		row.MarketValue = b.MarketValue
		row.TotalEquity = b.TotalEquity
		row.BuyingPower = b.BuyingPower
		row.MaintenanceExcess = b.MaintenanceExcess
		row.IsRealTime = b.IsRealTime
		csvB.Rows = append(csvB.Rows, row)
	}
}

func (csvB CSVBalance) ToSlice() []string {
	result := []string{
		csvB.Date.Format("01-02-2006 15:04:05"),
		csvB.AccountType,
		csvB.AccountNumber,
		csvB.BalanceType,
		csvB.Currency,
		fmt.Sprintf("%f", csvB.Cash),
		fmt.Sprintf("%f", csvB.MarketValue),
		fmt.Sprintf("%f", csvB.TotalEquity),
		fmt.Sprintf("%f", csvB.BuyingPower),
		fmt.Sprintf("%f", csvB.MaintenanceExcess),
		fmt.Sprintf("%t", csvB.IsRealTime),
	}
	return result
}

func (csvB *CSVBalances) WriteToCsv(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		writer := csv.NewWriter(f)
		err = writer.Write(CSVBalanceHeaders())
		if err != nil {
			log.Fatal("Cannot write headers to file", err)
		}
		writer.Flush()
		f.Close()
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot append to file", err)
	}
	writer := csv.NewWriter(f)
	for _, row := range csvB.Rows {
		err = writer.Write(row.ToSlice())
		if err != nil {
			log.Fatal("Cannot write row to file", err)
		}
	}
	writer.Flush()
	f.Close()
}
