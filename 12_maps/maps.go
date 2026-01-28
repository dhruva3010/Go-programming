package main

import "fmt"

func main() {
	// mapsExample()
	makeExample()
}

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println("Float Map:", f)
}

func mapsExample() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter":  "https://www.twitter.com",
	}

	fmt.Println("Websites Map:", websites)
	fmt.Println("Google:", websites["Google"])

	// append
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println("After adding LinkedIn:", websites)

	// delete
	delete(websites, "Twitter")
	fmt.Println("After deleting Twitter:", websites)
}

func makeExample() {
	// slices using make
	userNames := make([]string, 2, 5)

	userNames[0] = "Alice"
	userNames[1] = "Bob"

	userNames = append(userNames, "Charlie")
	fmt.Println("User Names Slice:", userNames)

	// maps using make
	courses := make(floatMap, 3)
	courses["Math"] = 95.5
	courses["Science"] = 89.0
	courses["History"] = 92.0

	courses.output()

	// for loop in maps
	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// for loop in maps
	for key, value := range courses {
		fmt.Printf("Course: %s, Score: %.2f\n", key, value)
	}
}
