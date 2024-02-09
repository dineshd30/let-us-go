package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Sentinel errors i.e. defined expected errors for better signalling
var (
	ErrInternal   = errors.New("internal error")
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)

func doSomething(value int) error {
	// some brilliant logic

	err := errors.New("some internal unknown error")
	// Wrap sentinel error with original error
	switch value {
	case 1:
		return fmt.Errorf("failed to get something: %w : %v", ErrNotFound, err)

	case 2:
		return fmt.Errorf("failed to get valid request: %w : %v", ErrBadRequest, err)

	case 3:
		return fmt.Errorf("failed to process: %w : %v", ErrInternal, err)

	case 4:
		return fmt.Errorf("unknown error: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Implementing error wrapping - Enter number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	err := doSomething(value)
	switch {
	case errors.Is(err, ErrBadRequest):
		fmt.Printf("Got http 400 bad request - %v\n", err) // return http 400 bad request

	case errors.Is(err, ErrNotFound):
		fmt.Printf("Got http 404 not found - %v\n", err) // return http 404 not found

	case errors.Is(err, ErrInternal):
		fmt.Printf("Got http 500 internal server error - %v\n", err) // return http 500 internal server error

	case err == nil:
		fmt.Println("No error")

	default:
		fmt.Printf("Got unexpected error - %v\n", err)
	}

}
