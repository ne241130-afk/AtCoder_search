package repository

import (
	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

type ProblemTagRepository struct{}

func (r *ProblemTagRepository) AttachTag(problemID string, tagID uint) error {
	return database.DB.Create(&model.ProblemTag{ProblemID: problemID, TagID: tagID}).Error
}

func (r *ProblemTagRepository) DetachTag(problemID string, tagID uint) error {
	return database.DB.Where("problem_id = ? AND tag_id = ?", problemID, tagID).Delete(&model.ProblemTag{}).Error
}
