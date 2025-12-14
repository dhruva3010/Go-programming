package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	print("Enter Total Revenue:")
	fmt.Scan(&revenue)
	print("Enter Total Expenses:")
	fmt.Scan(&expenses)
	print("Enter Tax Rate (in %):")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Earnings Before Tax (EBT):", ebt)
	fmt.Println("Net Profit:", profit)
	fmt.Println("EBT to Profit Ratio:", ratio)
}
