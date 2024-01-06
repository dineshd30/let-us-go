package main

import "fmt"

// Enum type
type Operation int

// Enum values
const (
	Add Operation = iota
	Sub
	Mul
)

func main() {
	fmt.Println("Demo : function literals")
	add := getCalculator(Add)
	fmt.Printf("Adding %d + %d = %d\n", 2, 4, add(2, 4))
}

// Function type calculator
type calculator func(int, int) int

func getCalculator(cmd Operation) calculator {
	var cal calculator
	switch cmd {
	case Add:
		cal = func(a int, b int) int {
			return a + b
		}

	case Sub:
		cal = func(a int, b int) int {
			return a - b
		}

	case Mul:
		cal = func(a int, b int) int {
			return a - b
		}
	}

	return func(a int, b int) int {
		return cal(a, b)
	}
}
