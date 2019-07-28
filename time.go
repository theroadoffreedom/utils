package utils

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"

	"github.com/go-shadow/moment"
)

const (
	DAY_TIMESTAMP_COUNT = 24 * 60 * 60
)

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func GetCurrentHourOf24() int {
	return moment.New().Hour()
}

func Translate2Timestamp(strDate string, format string) (int64, error) {
	t, err := time.Parse(format, strDate)
	if err != nil {
		return -1, err
	}
	return t.Unix(), nil
}

// return now(), format 2019-01-01 timestamp
func GetCurrentDayAlignTimestamp() int64 {
	alignStr := fmt.Sprintf("%s-%s-%sT00:00:00+08:00", GetCurrentYear(), GetCurrentMonthOnly(), GetCurrentDayOnly())
	tt, err := Translate2Timestamp(alignStr, time.RFC3339)
	if err != nil {
		fmt.Println(alignStr)
		fmt.Println(err)
	}
	return tt
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

func GetCurrentMonthOnly() string {
	m := moment.New().Month()	
	im := int32(m)
	if im < 10 {
		return fmt.Sprintf("0%d", im)
	}
	return fmt.Sprintf("%d",im)
}

func GetCurrentDayOnly() string {
	d := moment.New().DayOfMonth()
	id := int32(d)
	if id < 10 {
		return fmt.Sprintf("0%d", id)
	}
	return fmt.Sprintf("%d", id) 
}

// return current month, like 2019-07
func GetCurrentMonth() string {
	m := moment.New().Month()	
	y := GetCurrentYear()
	im := int32(m)
	if im < 10 {
		return fmt.Sprintf("%s-0%d", y, int32(m))
	}
	return fmt.Sprintf("%s-%d", y, int32(m))
}

// return next month, like 2019-08
func GetNextMonth(currentMonth string) (string, error) {

	str := strings.Split(currentMonth,"-")
	if len(str) != 2 {
		return "", errors.New("current month format is error")
	}
	mStr := str[1]
	yStr := str[0]

	mm := moment.New()
	y,err:=strconv.Atoi(yStr)
	if err != nil {
		return "",err
	}
	_mm := mm.SetYear(y)
	m, err := strconv.Atoi(mStr)
	if err != nil {
		return "", err
	}
	__mm := _mm.SetMonth(time.Month(m))
	___mm := __mm.AddMonths(1)
	ny := ___mm.Year()
	nm := int32(___mm.Month())

	if nm < 10 {
		return fmt.Sprintf("%d-0%d", ny, int32(nm)), nil
	}
	return fmt.Sprintf("%d-%d", ny, int32(nm)),nil
}

// return pass or not
func IsMonthPass(srcMonth string, dstMonth string) (bool, error) {
	// compare year
	str := strings.Split(srcMonth,"-")
	if len(str) != 2 {
		return false, errors.New("current month format is error")
	}
	mSrc,err:= strconv.Atoi(str[1])
	if err != nil {
		return false, err
	}
	ySrc,err:= strconv.Atoi(str[0])
	if err != nil {
		return false, err
	}
	str = strings.Split(dstMonth,"-")
	if len(str) != 2 {
		return false, errors.New("current month format is error")
	}
	mDst,err:= strconv.Atoi(str[1])
	if err != nil {
		return false, err
	}
	yDst,err:= strconv.Atoi(str[0])
	if err != nil {
		return false, err
	}

	if ySrc > yDst {
		return true,nil
	}
	if ySrc < yDst {
		return false,nil
	}

	// year equal
	if mSrc > mDst {
		return true, nil
	}
	return false,nil
}
