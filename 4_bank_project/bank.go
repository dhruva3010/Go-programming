package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func getBalanceFromFile() (data float64, err error) {
	content, err := os.ReadFile(balanceFileName)
	if err != nil {
		return 0.0, errors.New("could not read balance file")
	}
	balancetext := string(content)
	data, err = strconv.ParseFloat(balancetext, 64)
	if err != nil {
		return 0.0, errors.New("could not parse balance")
	}
	return
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(("-------------------------------"))
		return
		// panic("Could not initialize account balance")
	}

	fmt.Println("Welcome to the Bank Management System")

	for {
		fmt.Println("-------------------------------")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Deposit successful! New balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdrawal amount!")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawal successful! New balance is:", accountBalance)
				writeBalanceToFile(accountBalance)
			}
		case 4:
			fmt.Println("Exiting...")
			fmt.Println("Thank you for using the Bank Management System!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
