package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func WriteJSONToFile(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		file.Close()
		return err
	}
	file.Close()
	return nil
}
