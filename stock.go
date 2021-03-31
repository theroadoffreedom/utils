package utils

import (
	"fmt"
	"regexp"
	"strings"
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

func StockIDFormatCheck(id string) error {
	index := strings.Index(id, ".")
	if index != 2 {
		return fmt.Errorf("stock id format error")
	}
	return nil
}

// a股
// - sh
// - sz
// 美股
// - us
// 港股
// - hk
func GetMarketFromStockID(id string) (string, error) {

	err := StockIDFormatCheck(id)
	if err != nil {
		return "", err
	}

	if len(id) <= 2 {
		return "", fmt.Errorf("stock id error")
	}
	market := id[0:2]
	switch market {
	case "SH":
		{
			return market, nil
		}
	case "SZ":
		{
			return market, nil
		}

	case "HK":
		{
			return market, nil
		}
	case "US":
		{
			return market, nil
		}
	default:
		{
			return "", fmt.Errorf("not support stock id")
		}
	}
	return "", fmt.Errorf("not support stock id")
}

// SZ.000001
// hk.00700
// US.APPL
func GetStockCodeFromStockID(id string) (string, error) {

	err := StockIDFormatCheck(id)
	if err != nil {
		return "", err
	}

	if len(id) <= 3 {
		return "", fmt.Errorf("stock id error")
	}
	code := id[3:]
	if len(code) == 0 {
		return "", fmt.Errorf("get stock code error")
	}
	return code, nil
}

//
// SZ.000001 => 70001
// SH.600660 => 90660
// SZ.300001 => 77001
// SH.688158 => ?, not support
func StockIDToHSGTStockCode(stockID string) (string, error) {

	market, err := GetMarketFromStockID(stockID)
	if err != nil {
		return "", err
	}
	if strings.ToLower(market) == "hk" || strings.ToLower(market) == "us" {
		return "", fmt.Errorf("StockIDToHSGTStockCode not support this id %s", stockID)
	}

	code, err := GetStockCodeFromStockID(stockID)
	if err != nil {
		return "", err
	}

	// sz main board
	if code[0:2] == "000" {
		return "70" + code[3:], nil
	}

	// sh main board
	if code[0:2] == "600" {
		return "90" + code[3:], nil
	}

	// sz 300
	if code[0:2] == "300" {
		return "77" + code[3:], nil
	}

	return "", fmt.Errorf("StockIDToHSGTStockCode not support code:%s", stockID)

}

//
func GetStockIDFromHSGTStockCode(stockCode string) (string, error) {

	//
	if len(stockCode) != 5 {
		return "", fmt.Errorf("GetStockIDFromHSGTStockCode stock code error")
	}

	//
	if stockCode[0:1] == "77" {
		return "300" + stockCode[2:], nil
	}

	if stockCode[0:1] == "90" {
		return "600" + stockCode[2:], nil
	}

	if stockCode[0:1] == "70" {
		return "000" + stockCode[2:], nil
	}
	return "", fmt.Errorf("GetStockIDFromHSGTStockCode not support this code :%s", stockCode)
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
	return "", fmt.Errorf("stock id error")
}
