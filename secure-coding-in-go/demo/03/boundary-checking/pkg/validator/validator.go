package validator

import (
	"fmt"
	"strconv"
)

type IntegerValidator struct{}

func NewIntegerValidator() *IntegerValidator {
	return &IntegerValidator{}
}

func (i *IntegerValidator) Validate(value string) (int, error) {
	num, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("invalid input")
	}
	if num < 1 || num > 100 {
		return 0, fmt.Errorf("input out of range")
	}
	return num, nil
}
