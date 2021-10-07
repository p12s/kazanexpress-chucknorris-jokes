package internal

import (
	"fmt"
	"sync"
)

// RandomJoke @title Getting random joke
func RandomJoke() {
	joke, err := GetRandom()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(joke.Value)
}

// Dump @title Getting bulk of jokes by each category
func Dump(count int) {
	categories, err := GetCategories()
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	var wg sync.WaitGroup
	for _, category := range categories {
		wg.Add(1)
		go func(category string) {
			defer wg.Done()
			jokes, err := getJokesByCategory(category, count)
			if err != nil {
				fmt.Println("error:", err.Error())
			}
			SaveToFile(jokes, category)
		}(category)
	}
	wg.Wait()

	fmt.Printf("%d files created - %d categories, with %d jokes in each of the categories\n",
		len(categories), len(categories), count)
}

// getJokesByCategory @title Getting N jokes by category
func getJokesByCategory(category string, count int) ([]Joke, error) {
	var jokeList []Joke
	for i := 0; i < count; i++ {
		joke, err := getByCategory(category)
		if err != nil {
			return jokeList, err
		}
		jokeList = append(jokeList, joke)
	}
	return jokeList, nil
}

// getByCategory @title Getting one joke by category
func getByCategory(category string) (Joke, error) {
	return GetByCategory(category)
}

// Clear @title Clear all joke-saved files
func Clear() {
	categories, err := GetCategories()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	for _, category := range categories {
		clearFileIfExists(category)
	}
	fmt.Println("All files cleared")
}
