package main

import "fmt"

type transformFn func(int) int

func main() {
	funcAsPramExample()
	anonymousFuncExample()
	recursionExample()
	variadicExample()
}

func variadicExample() {
	fmt.Println("Variadic Function Example:")
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumUp(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", sum)

	spreadSum := sumUp(1, numbers...)
	fmt.Println("Sum of numbers using spread operator:", spreadSum)
}

func sumUp(first int, numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// recursion
func recursionExample() {
	num := 5
	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func anonymousFuncExample() {
	fmt.Println("Using Anonymous Functions as Parameters:")
	numbers := []int{1, 2, 3, 4, 5}
	double := transformerFunctionClosure(2)
	triple := transformerFunctionClosure(3)
	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Doubled Numbers:", transformNumbers(&numbers, double))
	fmt.Println("Tripled Numbers:", transformNumbers(&numbers, triple))
}

func transformerFunctionClosure(factor int) transformFn {
	return func(num int) int {
		return num * factor
	}
}

func funcAsPramExample() {
	numbers := []int{1, 2, 3, 4, 5}
	secNumbers := []int{10, 20, 30, 40, 50}
	doubledArray := transformNumbers(&numbers, double)
	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Doubled Numbers:", doubledArray)

	tripledArray := transformNumbers(&numbers, triple)
	fmt.Println("Tripled Numbers:", tripledArray)

	transforms1 := getTransformerFn(&numbers)
	transforms2 := getTransformerFn(&secNumbers)

	fmt.Println("Transform using function returned from another function:")
	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Second Numbers:", secNumbers)

	fmt.Println("Transformed Numbers 1:", transformNumbers(&numbers, transforms1))
	fmt.Println("Transformed Numbers 2:", transformNumbers(&secNumbers, transforms2))
}

// function as parameter
func transformNumbers(nums *[]int, transformFunc transformFn) []int {
	transformed := make([]int, len(*nums))
	for i, v := range *nums {
		transformed[i] = transformFunc(v)
	}
	return transformed
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func getTransformerFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
