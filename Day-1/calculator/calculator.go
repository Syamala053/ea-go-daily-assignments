package main

import "fmt"
import "math"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main() {
	fmt.Println("Calculator")
	fmt.Println("Addition of 2 and 4", add(2, 4))
	fmt.Println("Subtraction of 2 and 4", sub(2, 4))
	fmt.Println("Multiplication of 2 and 4", mul(2, 4))
	fmt.Println("Division of 2 and 4", div(2, 4))


	fmt.Println("Sin of 1 is ", math.Sin(1))
	fmt.Println("Cos of 1 is ", math.Cos(1))
	fmt.Println("Tan of 1 is ", math.Tan(1))

	fmt.Println("Square root of 2 is ", math.Sqrt(2))
}
