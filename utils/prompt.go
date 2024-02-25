package utils

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// Prompts

func Prompt(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: ValidateRequired,
	}

	return prompt.Run()
}

func PromptSelect(label string, items []string) (int, string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt.Run()
}

// Validation

func ValidateRequired(value string) error {
	if value == "" {
		return errors.New("May may not be empty")
	}

	return nil
}
