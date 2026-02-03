package gosugar

import (
	"errors"
	"fmt"
)

type Validator func(string) error

func NotEmpty() Validator {
	return func(s string) error {
		if s == "" {
			return errors.New("value cannot be empty")
		}
		return nil
	}
}

func MinLen(n int) Validator {
	return func(s string) error {
		if len(s) < n {
			return fmt.Errorf("minimum length is %d", n)
		}
		return nil
	}
}

func MaxLen(n int) Validator {
	return func(s string) error {
		if len(s) > n {
			return fmt.Errorf("maximum length is %d", n)
		}
		return nil
	}
}
