package internal

type Categories struct {
	Category []string `json:"category"`
}

type Joke struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconUrl    string   `json:"icon_url"`
	Id         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	Url        string   `json:"url"`
	Value      string   `json:"value"`
}
