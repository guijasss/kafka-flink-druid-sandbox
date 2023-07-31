package utils

import (
	"strconv"
)

func StringToFloat(floatString string) float32 {
	floatValue, _ := strconv.ParseFloat(floatString, 64)
	return float32(floatValue)
}
