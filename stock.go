package utils

import (
	"fmt"
	"regexp"
)

type StockPlate int

const (
	UnknowPlate StockPlate = iota
	SZ
	SH
	HK
	US
	JP
	K
)

type FinanceReportType int

const (
	AllSheet        FinanceReportType = iota
	BalanceSheet                      // 资产负债表
	CashStatement                     // 现金流表
	ProfitStatement                   // 利润表
)

type FinanceTimeType int

const (
	AllTime FinanceTimeType = iota
	Quarter                 // 季度
	Yearly                  // 年度
)

type ReportState int

const (
	AllReportState ReportState = iota
	ReportNormal
	ReportInvalid
	ReportFailed
)

// a股
// - sh
// - sz
// 美股
// - us
// 港股
// - hk
func GetMarketFromStockID(id string) (string, error) {
	if len(market) <= 2 {
		return "", fmt.Errorf("stock id error")
	}
	market := id[0:2]

	switch market {
	case "SH":
	case "SZ":
	case "HK":
	case "US":
		{
			return market, nil
		}
	default:
		{
			return "", fmt.Errorf("not support stock id")
		}
	}
}

// SZ.000001
// hk.00700
// US.APPL
func GetStockCodeFromStockID(id string) (string, error) {
	if len(market) <= 3 {
		return "", fmt.Errorf("stock id error")
	}
	return id[3:0], nil
}

//
// SZ.000001 => 70001
// SH.600660 => 90660
// SZ.300001 => 77001
// SH.688158 => ?, not support
func StockIDToHSGTStockCode(stockID string) (string, error) {

}

//
func GetStockIDFromHSGTStockCode(stockCode string) (string, error) {
}

func GetPlateStr(plate int) string {
	switch plate {
	case 0:
		{
			return "unknow"
		}
	case 1:
		{
			return "SH"
		}
	case 2:
		{
			return "SZ"
		}
	case 3:
		{
			return "HK"
		}
	case 4:
		{
			return "US"
		}
	default:
		{
			return "unknow"
		}
	}
}

func GetPlateFromStr(s string) int {
	switch s {
	case "unknow":
		{
			return 0
		}
	case "SH":
		{
			return 1
		}
	case "SZ":
		{
			return 2
		}
	case "HK":
		{
			return 3
		}
	case "US":
		{
			return 4
		}
	default:
		{
			return -1
		}
	}
}

func GetExchangeLabel(stock_id string) (string, error) {
	// macth 60****
	match, _ := regexp.MatchString("60[0-9][0-9][0-9][0-9]", stock_id)
	if match {
		return "SH", nil
	}

	// match 000***
	match, _ = regexp.MatchString("000[0-9][0-9][0-9]", stock_id)
	if match {
		return "SZ", nil
	}

	// match 001***
	match, _ = regexp.MatchString("001[0-9][0-9][0-9]", stock_id)
	if match {
		return "SZ", nil
	}

	// match 002***
	match, _ = regexp.MatchString("002[0-9][0-9][0-9]", stock_id)
	if match {
		return "SZ", nil
	}

	// match 30***
	match, _ = regexp.MatchString("30[0-9[0-9][0-9][0-9]", stock_id)
	if match {
		return "SZ", nil
	}

	// match 688**
	match, _ = regexp.MatchString("688[0-9][0-9][0-9]", stock_id)
	if match {
		return "SH", nil
	}

	// match 003**
	match, _ = regexp.MatchString("003[0-9][0-9][0-9]", stock_id)
	if match {
		return "SZ", nil
	}
	return "", errors.New("stock id error")
}
