package utils

import (
	"strconv"
	"time"
)

func StringConvertUint64(str string) uint64 {
	if num, err := strconv.ParseUint(str, 10, 64); err != nil {
		return 0
	} else {
		return num
	}
}

func StringConvertInt64(str string) int64 {
	if num, err := strconv.ParseInt(str, 10, 64); err != nil {
		return 0
	} else {
		return num
	}
}

func GetTimeNowUnix() int64 {
	return time.Now().Unix()
}
