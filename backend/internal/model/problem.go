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
}

type UserProblemStatus struct {
	ProblemID string `gorm:"primaryKey;column:problem_id" json:"problemId"`
	UserID    string `gorm:"index;not null;column:user_id" json:"userId"`
	Solved    bool   `gorm:"column:solved" json:"solved"`
	Favorite  bool   `gorm:"column:favorite" json:"favorite"`
	Memo      string `gorm:"column:memo" json:"memo"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at" json:"updatedAt"`
}

func (UserProblemStatus) TableName() string {
	return "user_problem_status"
}