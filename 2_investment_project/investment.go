package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Enter Expected Annual Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Enter Investment Duration (in years): ")
	fmt.Scan(&years)

	futureValue, futureRealValue := CalculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value of Investment: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value of Investment (Adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

	// Alternatively, using Printf for formatted output

	// fmt.Println("Future Value of Investment:", futureValue)
	// fmt.Printf("Future Value of Investment: %.2f\n", futureValue)

	// fmt.Println("Future Value of Investment (Adjusted for Inflation):", futureRealValue)
	// fmt.Printf("Future Value of Investment (Adjusted for Inflation): %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func CalculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}
