package repository

import (
	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

type TagRepository struct{}

func (r *TagRepository) FindAll() ([]model.Tag, error) {
	var tags []model.Tag
	err := database.DB.Order("name asc").Find(&tags).Error
	return tags, err
}

func (r *TagRepository) Create(name string) (*model.Tag, error) {
	tag := model.Tag{Name: name}
	if err := database.DB.Where("name = ?", name).FirstOrCreate(&tag, model.Tag{Name: name}).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) Delete(id uint) error {
	if err := database.DB.Where("tag_id = ?", id).Delete(&model.ProblemTag{}).Error; err != nil {
		return err
	}
	return database.DB.Delete(&model.Tag{}, id).Error
}
