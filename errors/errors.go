package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("Can't work with 2")
	}

	// A nil value in the error position indicates that there was no error.
	return arg, nil
}

// It’s possible to use custom types as errors by implementing the Error() method on them.
type customError struct {
	arg     int
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 2 {
		// In this case we use &customError syntax to build a new struct, supplying values for the two fields arg and prob.
		return -1, &customError{arg, "Cant't work with it"}
	}

	return arg, nil
}

func main() {
	for _, i := range []int{1, 2, 3, 4} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{1, 2, 3, 4} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	_, e := f2(2)
	// type assertion, ok = true so e is customError
	if ae, ok := e.(*customError); ok {
		fmt.Println(ok)
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	}
}
