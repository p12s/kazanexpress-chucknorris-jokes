package internal

// Categories @title Joke categories list
type Categories []string

// Joke @title Joke structure
type Joke struct {
	Categories Categories `json:"categories"`
	CreatedAt  string     `json:"created_at"`
	IconUrl    string     `json:"icon_url"`
	Id         string     `json:"id"`
	UpdatedAt  string     `json:"updated_at"`
	Url        string     `json:"url"`
	Value      string     `json:"value"`
}
