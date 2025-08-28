package main

import ( 
	"fmt"
)


func main(){

	// var result int
	var operation string
	var firstNumber float32
	var secondNumber float32
	var result float32
	

	fmt.Println("Enter operation:")
	fmt.Scanln(&operation)

	fmt.Println("Enter First Number:")
	fmt.Scanln(&firstNumber)

	fmt.Println("Enter Second Number:")
	fmt.Scanln(&secondNumber)


	// if ( operation == "*" ){
	// 	result = firstNumber * secondNumber
	// 	fmt.Println("first number",  firstNumber, "second number", secondNumber, "=", result)
	// }
		switch operation {
	case "*":
		result = firstNumber * secondNumber
		fmt.Println(firstNumber, operation, secondNumber, "=", result)
	case "+":
		result = firstNumber + secondNumber
		fmt.Println(firstNumber, operation, secondNumber, "=", result)
	case "-":
		result = firstNumber - secondNumber
		fmt.Println(firstNumber, operation, secondNumber, "=", result)
	case "/":
		result = firstNumber / secondNumber
		fmt.Println(firstNumber, operation, secondNumber, "=", result)
	default:
		fmt.Println("Unknown operation")
	}

	
}

