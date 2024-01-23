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

func someFunction(value int) error {
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
	err := someFunction(value)
	switch {
	case errors.Is(err, ErrBadRequest):
		fmt.Printf("Got ErrBadRequest - %v\n", err) // 400

	case errors.Is(err, ErrNotFound):
		fmt.Printf("Got ErrNotFound - %v\n", err) // 404

	case errors.Is(err, ErrInternal):
		fmt.Printf("Got ErrInternal - %v\n", err) // 500

	case err == nil:
		fmt.Println("No error")

	default:
		fmt.Printf("Got unexpected error - %v\n", err)
	}

}
