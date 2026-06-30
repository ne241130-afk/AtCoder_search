package model

import "github.com/lib/pq"

type Problem struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Contest     string         `gorm:"not null" json:"contest"`
	Difficulty  int            `gorm:"not null" json:"difficulty"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	AtCoderURL  string         `json:"atCoderUrl"`
	ProblemsURL string         `json:"problemsUrl"`
	Solved      bool           `json:"solved"`
	Favorite    bool           `json:"favorite"`
}