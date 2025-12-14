package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Enter Total Revenue: ")
	expenses := getUserInput("Enter Total Expenses: ")
	taxRate := getUserInput("Enter Tax Rate (in %): ")

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Println("Earnings Before Tax (EBT):", ebt)
	fmt.Println("Net Profit:", profit)
	fmt.Println("EBT to Profit Ratio:", ratio)
}

func getUserInput(text string) float64 {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

func calculateProfits(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
