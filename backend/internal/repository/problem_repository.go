package repository

import (
	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

type ProblemRepository struct{}

func (r *ProblemRepository) FindAll() ([]model.Problem, error) {
	var problems []model.Problem
	err := database.DB.Preload("Tags").Order("difficulty asc").Find(&problems).Error
	return problems, err
}

func (r *ProblemRepository) FindByID(id string) (*model.Problem, error) {
	var problem model.Problem
	err := database.DB.Preload("Tags").Where("id = ?", id).First(&problem).Error
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (r *ProblemRepository) Create(problem *model.Problem) error {
	return database.DB.Create(problem).Error
}
