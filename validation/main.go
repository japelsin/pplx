package validation

import (
	"errors"
	"strconv"
)

func ValidateInt(value string) error {
	_, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("Must be an int")
	}

	return nil
}

func ValidateRequired(value string) error {
	if value == "" {
		return errors.New("May not be empty")
	}

	return nil
}
