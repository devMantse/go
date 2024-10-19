package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find value from file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse storage")
	}

	return value, nil
}

func ValueToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
