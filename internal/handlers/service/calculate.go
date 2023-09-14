package service

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Сalculate(equation string) (int, error) {
	validPattern := "^[0-9+\\-]+$"
	match, err := regexp.MatchString(validPattern, equation)
	if err != nil {
		return 0, err
	}

	if !match {
		return 0, errors.New("некорректная строка")
	}

	operands := strings.Split(equation, "+")
	result := 0

	for _, operand := range operands {
		subOperands := strings.Split(operand, "-")
		subResult := 0
		for i, subOperand := range subOperands {
			num := 0
			_, err := fmt.Sscanf(subOperand, "%d", &num)
			if err != nil {
				return 0, err
			}
			if i == 0 {
				subResult += num
			} else {
				subResult -= num
			}
		}
		result += subResult
	}

	return result, nil
}
