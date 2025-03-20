package utils

import (
	"fmt"
	"strconv"
)

func StringToInt64(v string) int64 {
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 err:", err)
	}
	return val
}
