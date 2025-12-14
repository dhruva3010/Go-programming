package main

import (
	"fmt"

	"examples.com/bank_controls/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFileName)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(("-------------------------------"))
		// return
		// panic("Could not initialize account balance")
	}

	fmt.Println("Welcome to the Bank Management System")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
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
				fileops.WriteFloatToFile(balanceFileName, accountBalance)
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
