package internal

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	RANDOM_URL = "https://api.chucknorris.io/jokes/random"
)

var client = &http.Client{Timeout: 10 * time.Second}

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
