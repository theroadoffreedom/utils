package utils

import (
	//"fmt"
	"testing"
	//"time"
)

func TestGetMarketFromStockID(t *testing.T) {

	market, err := GetMarketFromStockID("SH.601012")
	if err != nil {
		t.Error(err)
	}
	t.Log(market)
}
