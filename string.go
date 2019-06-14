package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	UTILS_ERR_CODE_STRING_EMPTY              = "string is empty"
	UTILS_ERR_CODE_STRING_REMOVE_CHAR_FAILED = "remove char from string error"
)

func RemoveComma(s string) (string, error) {

	ss, err := RemoveStrings(s, ",")
	if err != nil {
		return "", errors.New(UTILS_ERR_CODE_STRING_REMOVE_CHAR_FAILED)
	}

	ss, err = RemoveStrings(ss, "，")
	return ss, err
}

func FinanceNumberStringToFloat(s string) (float32, error) {

	if len(s) == 0 {
		return 0, errors.New(UTILS_ERR_CODE_STRING_EMPTY)
	}

	ss, err := RemoveComma(s)
	if err != nil {
		return 0, err
	}

	f, err := String2Float(ss)
	return f, err
}
func FinanceNumberStringToInt(s string) (int, error) {

	i, err := FinanceNumberStringToInt32(s)
	if err != nil {
		return -1, err
	}
	return int(i), nil
}

func FinanceNumberStringToInt32(s string) (int32, error) {

	if len(s) == 0 {
		return -1, errors.New(UTILS_ERR_CODE_STRING_EMPTY)
	}

	ss, err := RemoveComma(s)
	if err != nil {
		return -1, err
	}

	i32, err := String2Int32(ss)
	return i32, err
}

func FinanceNumberStringToInt64(s string) (int64, error) {

	if len(s) == 0 {
		return -1, errors.New(UTILS_ERR_CODE_STRING_EMPTY)
	}

	ss, err := RemoveComma(s)
	if err != nil {
		return -1, err
	}

	i64, err := String2Int64(ss)
	return i64, err
}

func RemoveStrings(source string, chars string) (string, error) {

	s := strings.Split(source, chars)
	dest := ""
	for _, i := range s {
		dest = fmt.Sprintf("%s%s", dest, i)
	}
	return dest, nil
}

func String2Float(s string) (float32, error) {
	float, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(float), nil
}

func String2Int64(s string) (int64, error) {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, err
	}
	return i64, nil
}

func String2Int32(s string) (int32, error) {
	i32, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return int32(i32), nil
}
