package repository

import (
	"errors"
	"time"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
	"gorm.io/gorm"
)

type StatusRepository struct{}

func (r *StatusRepository) defaultUserID() string {
	return "default"
}

func (r *StatusRepository) GetStatus(problemID string) (*model.UserProblemStatus, error) {
	var status model.UserProblemStatus
	err := database.DB.Where("problem_id = ? AND user_id = ?", problemID, r.defaultUserID()).First(&status).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		status = model.UserProblemStatus{ProblemID: problemID, UserID: r.defaultUserID()}
		return &status, nil
	}
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusRepository) UpdateSolved(problemID string, solved bool) error {
	return r.upsert(problemID, func(status *model.UserProblemStatus) {
		status.Solved = solved
	})
}

func (r *StatusRepository) UpdateFavorite(problemID string, favorite bool) error {
	return r.upsert(problemID, func(status *model.UserProblemStatus) {
		status.Favorite = favorite
	})
}

func (r *StatusRepository) UpdateMemo(problemID string, memo string) error {
	return r.upsert(problemID, func(status *model.UserProblemStatus) {
		status.Memo = memo
	})
}

func (r *StatusRepository) upsert(problemID string, mutate func(*model.UserProblemStatus)) error {
	status, err := r.GetStatus(problemID)
	if err != nil {
		return err
	}

	mutate(status)
	status.UpdatedAt = time.Now().UnixMilli()

	return database.DB.Save(status).Error
}
