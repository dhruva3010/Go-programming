package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples.com/notes/note"
)

func main() {
	title, content := getNotesData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
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
