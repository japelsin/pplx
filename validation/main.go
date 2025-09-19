package validation

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/japelsin/pplx/constants"
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

func ValidateRecencyFilter(value string) error {
	isValidRecencyFilter := slices.Contains(constants.SEARCH_RECENCY_FILTERS, value)
	if !isValidRecencyFilter {
		return errors.New("Invalid recency filter, must be one of: " + strings.Join(constants.SEARCH_RECENCY_FILTERS[:], ", "))
	}

	return nil
}
