package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Expected Annual Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter Investment Duration (in years): ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value of Investment: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value of Investment (Adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

	// Alternatively, using Printf for formatted output

	// fmt.Println("Future Value of Investment:", futureValue)
	// fmt.Printf("Future Value of Investment: %.2f\n", futureValue)

	// fmt.Println("Future Value of Investment (Adjusted for Inflation):", futureRealValue)
	// fmt.Printf("Future Value of Investment (Adjusted for Inflation): %.2f\n", futureRealValue)
}
