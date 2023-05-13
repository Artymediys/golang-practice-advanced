package main

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorSlice []error

func (errs ErrorSlice) Error() string {
	stringErrors := make([]string, len(errs))
	for index, err := range errs {
		stringErrors[index] = err.Error()
	}

	return strings.Join(stringErrors, "; ")
}

func MyCheck(input string) error {
	var finalErr ErrorSlice

	if hasDigit := strings.ContainsAny(input, "0123456789"); hasDigit {
		finalErr = append(finalErr, errors.New("found digits"))
	}

	if len([]rune(input)) > 20 {
		finalErr = append(finalErr, errors.New("line is too long"))
	}

	var spaces uint8
	for _, char := range input {
		if char == ' ' {
			spaces++
		}
	}

	if spaces != 2 {
		finalErr = append(finalErr, errors.New("no two spaces"))
	}

	return finalErr
}

func main() {
	testString := "Test1ng some sussy stuff"

	err := MyCheck(testString)
	if err != nil {
		fmt.Println(err)
	}
}
