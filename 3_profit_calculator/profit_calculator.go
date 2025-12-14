package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, errRevenue := getUserInput("Enter Total Revenue: ")
	if errRevenue != nil {
		fmt.Println("Error:", errRevenue)
		return
	}
	expenses, errExpenses := getUserInput("Enter Total Expenses: ")
	if errExpenses != nil {
		fmt.Println("Error:", errExpenses)
		return
	}
	taxRate, errTax := getUserInput("Enter Tax Rate (in %): ")
	if errTax != nil {
		fmt.Println("Error:", errTax)
		return
	}

	// if errRevenue != nil || errExpenses != nil || errTax != nil {
	// 	fmt.Println("Invalid input. Please enter positive numbers only.")
	// 	return
	// }

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Println("Earnings Before Tax (EBT):", ebt)
	fmt.Println("Net Profit:", profit)
	fmt.Println("EBT to Profit Ratio:", ratio)
	storeToFile(ebt, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("input must be greater than zero")
	}

	return input, nil
}

func calculateProfits(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func storeToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nTax Rate: %.2f\n", ebt, profit, ratio)
	err := os.WriteFile("financial_data.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}
}
