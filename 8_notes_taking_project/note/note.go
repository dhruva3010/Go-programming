package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// with tags for json marshalling
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		fmt.Println("Input cannot be empty. Please try again.")
		return &Note{}, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Content:", n.Content)
	fmt.Println("Created At:", n.CreatedAt)
}

func (n *Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_") + ".json"
	filename = strings.ToLower(filename)

	data, err := json.Marshal(n) // Only publically available data in struct will be marshalled

	if err != nil {
		fmt.Println("Error marshalling note:", err)
		return err
	}

	perm := os.FileMode(0644)
	return os.WriteFile(filename, data, perm)
}
