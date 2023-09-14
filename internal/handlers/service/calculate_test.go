package service

import (
	"errors"
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
		err      error
	}{
		{input: "2+3-1", expected: 4, err: nil},
		{input: "10-5+3", expected: 8, err: nil},
		{input: "1+2+3+4+5", expected: 15, err: nil},

		{input: "", expected: 0, err: errors.New("некорректная строка")},

		{input: "abc+def", expected: 0, err: errors.New("некорректная строка")},
		{input: "1+2+3+4a5", expected: 0, err: errors.New("некорректная строка")},
		{input: "1++2", expected: 0, err: errors.New("некорректная строка")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result, err := Сalculate(testCase.input)

			if result != testCase.expected {
				t.Errorf("Expected result %d, but got %d", testCase.expected, result)
			}

			if err != nil {
				if testCase.err == nil {
					t.Errorf("Expected no error, but got %v", err)
				} else if reflect.TypeOf(err) != reflect.TypeOf(testCase.err) {
					t.Errorf("Expected error type %T, but got %T", testCase.err, err)
				}
			} else if testCase.err != nil {
				t.Errorf("Expected error %v, but got no error", testCase.err)
			}
		})
	}
}
