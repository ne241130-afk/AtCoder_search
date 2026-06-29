package model

type Problem struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Contest    string   `json:"contest"`
	Difficulty int      `json:"difficulty"`
	Tags       []string `json:"tags"`
}