package gobyexample

import (
	"errors"
	"fmt"
)

type CustomError struct {
	arg     int
	message string
}

func (err CustomError) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.message)
}

var ErrRandom = errors.New("Some error message")

func wrapError() error {
	return fmt.Errorf("wrapping error: %w", ErrRandom)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, CustomError{arg, "it is not working"}
	}
	return arg + 3, nil
}

func CustomErrors() {
	// errors.Is expects a target error (value) not a type
	if err := wrapError(); errors.Is(err, ErrRandom) {
		fmt.Println("A random error: ", err)
	}

	fmt.Println("===========")
	if _, err := f(42); err != nil {
		fmt.Println(err.Error())
	}

	_, err := f(42)
	
	// errors.AsType expects a type
	// errors.AsType checks if any error in the err's error tree matches a CustomError type
	if ae, ok := errors.AsType[CustomError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("is not of type CustomError")
	}
}
