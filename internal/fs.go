package internal

import (
	"fmt"
	"os"
)

func SaveToFile(jokes []Joke, categoryName string) {
	clearFileIfExists(categoryName)
	appendToFile(jokes, categoryName)
}

func clearFileIfExists(fileName string) {
	err := os.RemoveAll(fileName + ".txt")
	if err != nil {
		fmt.Println("Clear file error:", err.Error())
	}
}

func appendToFile(jokes []Joke, fileName string) {
	f, err := os.OpenFile(fileName + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Append to file error:", err.Error())
	}
	defer f.Close()

	for _, joke := range jokes {
		if _, err := f.WriteString(joke.Value + "\n"); err != nil {
			fmt.Println("Append to file error:", err.Error())
		}
	}
}
