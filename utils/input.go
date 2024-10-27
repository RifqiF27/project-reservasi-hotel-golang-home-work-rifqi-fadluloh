package utils

import (
	"errors"
	"fmt"
	// "regexp"
	"strconv"
)

func InputStr(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", errors.New("input invalid")
	}
	// isString := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	// if !isString(input) {
	// 	return "", errors.New("input must be letters")
	// }
	return input, nil
}
func InputInt(prompt string) (int, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("input invalid")
	} else if value <= 0 {
		return 0, errors.New("input must be greater than 0")
	}
	return value, nil
}

func InputBool(prompt string) (bool, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false, errors.New("input invalid")
	}

	value, err := strconv.ParseBool(input)

	if err != nil {
		if input == "available" {
			return true, nil
		} else if input == "unavailable" {
			return false, nil
		}
		return false, errors.New("input must be 'true', 'false', 'available', or 'unavailable'")
	}

	return value, nil
}
