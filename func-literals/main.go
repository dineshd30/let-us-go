package main

import "fmt"

func main() {
	fmt.Println("Demo : function literals")
	add := getCalculator("add")
	fmt.Printf("Adding %d + %d = %d\n", 2, 4, add(2, 4))
}

type calculator func(int, int) int

func getCalculator(cmd string) calculator {
	var operation calculator
	switch cmd {
	case "add":
		operation = func(a int, b int) int {
			return a + b
		}

	case "sub":
		operation = func(a int, b int) int {
			return a - b
		}

	case "mul":
		operation = func(a int, b int) int {
			return a - b
		}
	}

	return func(a int, b int) int {
		return operation(a, b)
	}
}
