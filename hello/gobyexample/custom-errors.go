package gobyexample

import (
	"errors"
	"fmt"
)

var ErrRandom = errors.New("Some error message")

func wrapError() error {
	return fmt.Errorf("wrapping error: %w", ErrRandom)
}

func CustomErrors() {
	if err := wrapError(); errors.Is(err, ErrRandom) {
		fmt.Println("A random error: ", err)
	}
}
