package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Position struct {
	Symbol             string  `json:"symbol"`
	SymbolId           int     `json:"symbolId"`
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

type CSVPositions struct {
	Rows []CSVPosition
}

type CSVPosition struct {
	Date               time.Time
	AccountType        string
	AccountNumber      string
	Symbol             string
	SymbolId           int
	OpenQuantity       float64
	ClosedQuantity     float64
	CurrentMarketValue float64
	CurrentPrice       float64
	AverageEntryPrice  float64
	DayPnl             float64
	ClosedPnl          float64
	OpenPnl            float64
	TotalCost          float64
	IsRealTime         bool
	IsUnderReorg       bool
}

func CSVPositionHeaders() []string {
	return []string{
		"Date",
		"AccountType",
		"AccountNumber",
		"Symbol",
		"SymbolId",
		"OpenQuantity",
		"ClosedQuantity",
		"CurrentMarketValue",
		"CurrentPrice",
		"AverageEntryPrice",
		"DayPnl",
		"ClosedPnl",
		"OpenPnl",
		"TotalCost",
		"IsRealTime",
		"IsUnderReorg",
	}
}

func (csvP *CSVPositions) FromPositions(positions []Position, account Account) {
	for _, p := range positions {
		var row CSVPosition
		row.Date = time.Now()
		row.AccountType = account.Type
		row.AccountNumber = account.Number
		row.Symbol = p.Symbol
		row.SymbolId = p.SymbolId
		row.OpenQuantity = p.OpenQuantity
		row.ClosedQuantity = p.ClosedQuantity
		row.CurrentMarketValue = p.CurrentMarketValue
		row.CurrentPrice = p.CurrentPrice
		row.AverageEntryPrice = p.AverageEntryPrice
		row.DayPnl = p.DayPnl
		row.ClosedPnl = p.ClosedPnl
		row.OpenPnl = p.OpenPnl
		row.TotalCost = p.TotalCost
		row.IsRealTime = p.IsRealTime
		row.IsUnderReorg = p.IsUnderReorg
		csvP.Rows = append(csvP.Rows, row)
	}
}

func (csvP CSVPosition) ToSlice() []string {
	result := []string{
		csvP.Date.Format("01-02-2006 15:04:05"),
		csvP.AccountType,
		csvP.AccountNumber,
		csvP.Symbol,
		fmt.Sprintf("%d", csvP.SymbolId),
		fmt.Sprintf("%f", csvP.OpenQuantity),
		fmt.Sprintf("%f", csvP.ClosedQuantity),
		fmt.Sprintf("%f", csvP.CurrentMarketValue),
		fmt.Sprintf("%f", csvP.CurrentPrice),
		fmt.Sprintf("%f", csvP.AverageEntryPrice),
		fmt.Sprintf("%f", csvP.DayPnl),
		fmt.Sprintf("%f", csvP.ClosedPnl),
		fmt.Sprintf("%f", csvP.OpenPnl),
		fmt.Sprintf("%f", csvP.TotalCost),
		fmt.Sprintf("%t", csvP.IsRealTime),
		fmt.Sprintf("%t", csvP.IsUnderReorg),
	}
	return result
}

func (csvP *CSVPositions) WriteToCsv(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		writer := csv.NewWriter(f)
		err = writer.Write(CSVPositionHeaders())
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
	for _, row := range csvP.Rows {
		err = writer.Write(row.ToSlice())
		if err != nil {
			log.Fatal("Cannot write row to file", err)
		}
	}
	writer.Flush()
	f.Close()
}
