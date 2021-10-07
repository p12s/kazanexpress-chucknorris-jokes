package internal

import "fmt"

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
	for _, category := range categories {
		jokes, err := getJokesByCategory(category, count)
		if err != nil {
			fmt.Println("error:", err.Error())
			break
		}
		SaveToFile(jokes, category)
	}
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