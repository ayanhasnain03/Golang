package main

import "fmt"

func main() {
	var num1, num2 float64
	var operation string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero.")
		}
	default:
		fmt.Println("Invalid operation.")
	}
}
