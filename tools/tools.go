package tools

import (
	"context"
	"strconv"
)

type (
	AppContex struct {
		Context context.Context
	}
)

func stringToFloat(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func floatToString(val float64) string {
	return strconv.FormatFloat(val, 'g', -1, 64)
}
