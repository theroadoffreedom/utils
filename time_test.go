package utils

import (
	"fmt"
	"testing"
	//	"time"
)

func TestGetYearFromTimestamp(t *testing.T) {

	tt, err := Translate2Timestamp("2018-10-19 15:04:05", "2006-01-02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tt)
	fmt.Println(GetYearFromTimestamp(tt))
	fmt.Println(GetMonthFromTimestamp(tt))
	fmt.Println(GetDayFromTimestamp(tt))
}


func TestGetCurrentMonth(t *testing.T) {
	mStr := GetCurrentMonth()
	t.Log(mStr)
}

func TestGetNextMonth(t *testing.T) {

	mStr := GetCurrentMonth()
	for i:= 0; i<15; i++ {
		m, err := GetNextMonth(mStr)
		if err != nil {
			t.Error(err.Error())
			return
		}
		t.Log(m)
		mStr = m
	}
}
