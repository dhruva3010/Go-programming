package conversions

import (
	"strconv"
)

func StringToFloat64(s []string) ([]float64, error) {
	prices := make([]float64, len(s))
	for lineIndex, line := range s {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		prices[lineIndex] = price
	}

	return prices, nil
}
