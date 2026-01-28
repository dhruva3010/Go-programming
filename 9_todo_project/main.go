package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples.com/notes/note"
	"examples.com/todos/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	saver
// 	displayer
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNotesData()

	todoText := getUserInput("Enter Todo Item: ")
	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		fmt.Println("Error outputting todo:", err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println("Error outputting note:", err)
		return
	}

	printSomething("Hello, World!")
	printSomething(42)
	printSomething(true)
	printSomething(3.14)
	printSomething(userTodo)
}

// interface usage
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	if err == nil {
		return err
	}
	return nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func getNotesData() (string, string) {
	title := getUserInput("Enter Note Title: ")

	content := getUserInput("Enter Note Content: ")

	return title, content
}

// switch types
func printSomething(value interface{}) { // also can use any
	intValue, ok := value.(int)

	if ok {
		fmt.Println("It's an integer:", intValue)
		return
	}

	stringValue, ok := value.(string)

	if ok {
		fmt.Println("It's a string:", stringValue)
		return
	}

	boolValue, ok := value.(bool)

	if ok {
		fmt.Println("It's a boolean:", boolValue)
		return
	}

	floatValue, ok := value.(float64)

	if ok {
		fmt.Println("It's a float64:", floatValue)
		return
	}

	fmt.Println("Unknown type")

	// switch value.(type) {
	// case string:
	// 	fmt.Println("It's a string:", value)
	// case int:
	// 	fmt.Println("It's an integer:", value)
	// case bool:
	// 	fmt.Println("It's a boolean:", value)
	// case float64:
	// 	fmt.Println("It's a float64:", value)
	// default:
	// 	fmt.Println("Unknown type")
	// }
}
