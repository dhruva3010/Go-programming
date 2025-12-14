package main

import "fmt"

func main() {
	age := 32
	agePtr := &age

	fmt.Println("Age Address:", agePtr)
	fmt.Println("Age:", *agePtr)

	EditAdultYears(agePtr)

	fmt.Println("Adult Years:", age)
}

// func getAAdultYears(age *int) int {
// 	return *age - 18
// }

// mutate value using pointer
func EditAdultYears(age *int) {
	*age = *age - 18
}
