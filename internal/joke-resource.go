package internal

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	URL            = "https://api.chucknorris.io/jokes/"
	RANDOM_URL     = URL + "random"
	CATEGORIES_URL = URL + "categories"
	CATEGORY_URL   = URL + "random?category="
)

// client @title Client for HTTP-request
var client = &http.Client{Timeout: 10 * time.Second}

// GetRandom @title get random joke without category picking
func GetRandom() (Joke, error) {
	var joke Joke
	res, err := client.Get(RANDOM_URL)
	if err != nil {
		return joke, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&joke)
	return joke, err
}

// GetByCategory @title Get random joke by category
func GetByCategory(category string) (Joke, error) {
	var joke Joke
	res, err := client.Get(CATEGORY_URL + category)
	if err != nil {
		return joke, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&joke)
	return joke, err
}

// GetCategories @title Getting categories list
func GetCategories() (Categories, error) {
	var categories Categories
	res, err := client.Get(CATEGORIES_URL)
	if err != nil {
		return categories, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&categories)
	return categories, err
}
