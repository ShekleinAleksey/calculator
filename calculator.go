package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// add := flag.Bool("add", false, "Add two numbers")
	// subs := flag.Bool("subtract", false, "Subtract two numbers")
	// mult := flag.Bool("multiply", false, "Multiply two numbers")
	// div := flag.Bool("divide", false, "Divide two numbers")
	var operator string
	fmt.Println("Enter operator: +, -, *, /")
	fmt.Scan(&operator)

	flag.Parse()
	var firstNum, secondNum float64
	fmt.Println("1 число:")
	fmt.Scan(&firstNum)
	fmt.Println("2 число:")
	fmt.Scan(&secondNum)
	var result float64
	switch operator {
	case "+":
		result = addition(firstNum, secondNum)
	case "-":
		result = substract(firstNum, secondNum)
	case "*":
		result = multiply(firstNum, secondNum)
	case "/":
		if secondNum == 0 {
			log.Fatal("Cannot divide by zero")
		}
		result = division(firstNum, secondNum)
	default:
		fmt.Fprintln(os.Stderr, "Wrong option try with add, subtract, div and multply")
		os.Exit(1)
	}
	fmt.Printf("Result: %.0f %s %.0f = %.0f", firstNum, operator, secondNum, result)
}

func addition(a, b float64) float64 {
	return a + b
}

func substract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	return a / b
}
