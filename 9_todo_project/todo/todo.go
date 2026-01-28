package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// with tags for json marshalling
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("Todo:", todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	data, err := json.Marshal(todo) // Only publically available data in struct will be marshalled

	if err != nil {
		fmt.Println("Error marshalling todo:", err)
		return err
	}

	perm := os.FileMode(0644)
	return os.WriteFile(filename, data, perm)
}

func New(content string) (Todo, error) {
	if content == "" {
		fmt.Println("Input cannot be empty. Please try again.")
		return Todo{}, errors.New("content cannot be empty")
	}
	return Todo{
		Text: content,
	}, nil
}
