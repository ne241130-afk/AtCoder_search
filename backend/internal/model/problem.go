package model

type Problem struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Contest     string `gorm:"not null" json:"contest"`
	Difficulty  int    `gorm:"not null" json:"difficulty"`
	AtCoderURL  string `json:"atCoderUrl"`
	ProblemsURL string `json:"problemsUrl"`
	Tags        []Tag  `gorm:"many2many:problem_tags;joinForeignKey:problem_id;joinReferences:tag_id" json:"tags"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

type ProblemTag struct {
	ProblemID string `gorm:"primaryKey;column:problem_id" json:"problemId"`
	TagID     uint   `gorm:"primaryKey;column:tag_id" json:"tagId"`
}

func (ProblemTag) TableName() string {
	return "problem_tags"
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
