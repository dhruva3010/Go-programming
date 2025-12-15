package main

import (
	"fmt"

	"examples.com/structs/user"
)

func main() {
	var firstName, lastName, birthDate string

	firstName = getUserInput("Enter First Name: ", firstName)
	lastName = getUserInput("Enter Last Name: ", lastName)
	birthDate = getUserInput("Enter Birth Date (YYYY-MM-DD): ", birthDate)

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()

	adminUser, err := user.NewAdmin("admin@example.com", "password123", "Admin", "Test", "1990-01-01")

	if err != nil {
		fmt.Println("Error creating admin user:", err)
		return
	}

	adminUser.OutputUserDetails()
}

// func outputUserDetails(user *User) {
// 	// no need for dereferencing pointer to access fields of struct
// 	fmt.Println("User Details:")
// 	fmt.Println("First Name:", user.firstName)
// 	fmt.Println("Last Name:", user.lastName)
// 	fmt.Println("Birth Date:", user.birthDate)
// 	fmt.Println("Account Created At:", user.createdAt)
// }

func getUserInput(prompt, input string) string {
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
