package utils

import (
	"fmt"
	"time"
)

const (
	DAY_TIMESTAMP_COUNT = 24 * 60 * 60
)

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func Translate2Timestamp(strDate string, format string) (int64, error) {
	t, err := time.Parse(format, strDate)
	if err != nil {
		return -1, err
	}
	return t.Unix(), nil
}

func AlignDailyTimestamp(t int64) int64 {

	//return t - (t % int64(DAY_TIMESTAMP_COUNT))
	alignStr := fmt.Sprintf("%s-%s-%sT00:00:00+08:00", GetYearFromTimestamp(t), GetMonthFromTimestamp(t), GetDayFromTimestamp(t))
	tt, err := Translate2Timestamp(alignStr, time.RFC3339)
	if err != nil {
		fmt.Println(alignStr)
		fmt.Println(err)
	}
	return tt
}

func ToHumanString(t int64) string {
	tt := time.Unix(t, 0)
	return tt.String()
}

// return string year
func GetYearFromTimestamp(t int64) string {
	return fmt.Sprintf("%d", time.Unix(t, 0).Year())
}

// return string month
func GetMonthFromTimestamp(t int64) string {
	return fmt.Sprintf("%d", int(time.Unix(t, 0).Month()))
}

// return string day
func GetDayFromTimestamp(t int64) string {
	return fmt.Sprintf("%d", time.Unix(t, 0).Day())
}

// return current string year
func GetCurrentYear() string {
	return GetYearFromTimestamp(GetCurrentTimestamp())
}
