package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	staticArrayExample()
	dynamicArrayExample()
	practiceFunction()
	appendSlices()
}

func staticArrayExample() {
	fmt.Println("Static Arrays and Slices Example")
	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
	var productNames = [3]string{"Laptop", "Smartphone", "Tablet"}

	fmt.Println("Product Names Array:", productNames)

	fmt.Println("Prices Array:", prices)

	// array slice
	featuredPrices := prices[1:3]
	fmt.Println("Featured Prices Slice:", featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println("Featured Prices Slice (from start):", featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println("Featured Prices Slice (from index 1):", featuredPrices)

	highlightedPrices := featuredPrices[:1]

	fmt.Println("Lenght of array prices:", len(featuredPrices), "capacity of array prices:", cap(featuredPrices))
	fmt.Println("Highlighted Prices Slice:", highlightedPrices)
	fmt.Println("Lenght of array prices:", len(highlightedPrices), "capacity of array prices:", cap(highlightedPrices))

	// can reslice to the right within capacity
	extendedPrices := highlightedPrices[:3]
	fmt.Println("Extended Prices Slice:", extendedPrices)
}

func dynamicArrayExample() {
	fmt.Println("\nDynamic Arrays (Slices) Example")

	prices := []float64{10.99, 20.99, 30.99}

	fmt.Println("Initial Prices Slice:", prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println("Prices Slice after append:", updatedPrices)
}

func practiceFunction() {
	fmt.Println("\nPractice Function Example")

	// 1. create array using 3 hobbies
	hobbies := [3]string{"Reading", "Traveling", "Cooking"}
	fmt.Println("Hobbies:", hobbies)

	// 2. outputs 1st element standalone and 2nd and 3rd as slice
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("Other hobbies:", hobbies[1:3])

	// 3. create a slice with 1st and 2nd element
	mainHobbies := hobbies[:2]
	fmt.Println("Main Hobbies Slice:", mainHobbies)

	// 4. reslice from 3rd point and display 2nd and lats element in the array
	otherHobbies := mainHobbies[1:3] // need to specify extended slice to access 3rd element
	fmt.Println("Other Hobbies Slice:", otherHobbies)

	// 5. create a dynamic array for course goals
	courseGoals := []string{"Learn Go", "Build Projects"}
	fmt.Println("Initial Course Goals Slice:", courseGoals)

	// 6. update 2nd goal and append a new goal
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Learn all the basics")

	fmt.Println("Updated Course Goals Slice:", courseGoals)

	// 7. create product struct array
	products := []Product{
		{id: 1, title: "Book", price: 9.99},
		{id: 2, title: "Pen", price: 1.99},
	}
	fmt.Println("Products Slice:", products)

	products = append(products, Product{id: 3, title: "Notebook", price: 4.99})
	fmt.Println("Products Slice after append:", products)
}

func appendSlices() {
	fmt.Println("\nAppend Slices Example")
	prices := []float64{10.99, 20.99, 30.99}

	fmt.Println("Initial Prices Slice:", prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println("Prices Slice after append:", updatedPrices)

	discountedPrices := []float64{8.99, 18.99}
	combinedPrices := append(prices, discountedPrices...) // rest operator to get all elements from discountedPrices
	fmt.Println("Combined Prices Slice:", combinedPrices)
}
