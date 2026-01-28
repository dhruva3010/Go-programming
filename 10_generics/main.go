package main

import "fmt"

func main() {
	result := add(5, 10)
	fmt.Println("Sum of integers:", result)

	resultFloat := add(5.5, 10.3)
	fmt.Println("Sum of floats:", resultFloat)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
