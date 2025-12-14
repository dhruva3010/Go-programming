package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Use caps at the start of function name to make it public
func GetFloatFromFile(fileName string) (data float64, err error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("could not read file")
	}
	valueText := string(content)
	data, err = strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0.0, errors.New("could not parse va;ue")
	}
	return
}

func WriteFloatToFile(fileName string, value float64) error {
	valueText := fmt.Sprint(value)
	return os.WriteFile(fileName, []byte(valueText), 0644)
}
