package internal

import (
	"encoding/json"
	"net/http"
)

const (
	URL            = "https://api.chucknorris.io/jokes/"
	RANDOM_URL     = URL + "random"
	CATEGORIES_URL = URL + "categories"
	CATEGORY_URL   = URL + "random?category="
)

// GetRandom @title get random joke without category picking
func GetRandom() (Joke, error) {
	var joke Joke
	res, err := http.Get(RANDOM_URL)
	if err != nil {
		return joke, err
	}

	err = json.NewDecoder(res.Body).Decode(&joke)
	if err != nil {
		_ = res.Body.Close()
		return joke, err
	}

	closeErr := res.Body.Close()
	return joke, closeErr
}

// GetByCategory @title Get random joke by category
func GetByCategory(category string) (Joke, error) {
	var joke Joke
	res, err := http.Get(CATEGORY_URL + category)
	if err != nil {
		return joke, err
	}

	err = json.NewDecoder(res.Body).Decode(&joke)
	if err != nil {
		_ = res.Body.Close()
		return joke, err
	}

	closeErr := res.Body.Close()
	return joke, closeErr
}

// GetCategories @title Getting categories list
func GetCategories() (Categories, error) {
	var categories Categories
	res, err := http.Get(CATEGORIES_URL)
	if err != nil {
		return categories, err
	}

	err = json.NewDecoder(res.Body).Decode(&categories)
	if err != nil {
		_ = res.Body.Close()
		return categories, err
	}

	closeErr := res.Body.Close()
	return categories, closeErr
}
