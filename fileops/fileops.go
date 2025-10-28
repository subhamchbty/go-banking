package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(value float64, filename string) error {
	byteText := []byte(fmt.Sprint(value))
	err := os.WriteFile(filename, byteText, 0644)

	return err
}

func ReadValueFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0, errors.New("cannot get value from file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("cannot convert string to float")
	}

	return value, nil
}
