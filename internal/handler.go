package internal

import (
	"fmt"
	"sync"
)

// @title User-comment API
// @version 0.1
// @description API Server for users and their comments
// @host localhost:80
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RandomJoke() {
	joke, err := GetRandom()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(joke.Value)
}

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

func getByCategory(category string) (Joke, error) {
	return GetByCategory(category)
}

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