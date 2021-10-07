package internal

import (
	"fmt"
	"os"
)

// SaveToFile @title Save one category jokes to file
func SaveToFile(jokes []Joke, categoryName string) {
	clearFileIfExists(categoryName)
	appendToFile(jokes, categoryName)
}

// clearFileIfExists @title Remove file if exists
func clearFileIfExists(fileName string) {
	err := os.RemoveAll(fileName + ".txt")
	if err != nil {
		fmt.Println("Clear file error:", err.Error())
	}
}

// appendToFile @title Append joke to file
func appendToFile(jokes []Joke, fileName string) {
	f, err := os.OpenFile(fileName+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Append to file error:", err.Error())
	}

	for _, joke := range jokes {
		if _, err := f.WriteString(joke.Value + "\n"); err != nil {
			fmt.Println("Append to file error:", err.Error())
		}
	}

	closeErr := f.Close()
	if closeErr != nil {
		fmt.Println("Close file error:", closeErr.Error())
	}
}
