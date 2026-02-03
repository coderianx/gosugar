package gosugar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputRaw(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		panic("input error")
	}

	return strings.TrimSpace(scanner.Text())
}

func Input(prompt string, validators ...Validator) string {
	input := inputRaw(prompt)

	for _, validate := range validators {
		if err := validate(input); err != nil {
			panic(fmt.Errorf("invalid string input: %w", err))
		}
	}

	return input
}

func InputInt(prompt string, defaultValue ...int) int {
	input := inputRaw(prompt)

	value, err := strconv.Atoi(input)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("invalid integer input: %q", input))
	}

	return value
}

func InputFloat(prompt string, defaultValue ...float64) float64 {
	input := inputRaw(prompt)

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("invalid float input: %q", input))
	}

	return value
}
