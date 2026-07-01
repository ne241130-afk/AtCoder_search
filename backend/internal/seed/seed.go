package seed

import (
	"log"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
	"gorm.io/gorm/clause"
)

func SeedProblems() error {
	var count int64
	if err := database.DB.Model(&model.Problem{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		for _, problem := range Problems {
			problemCopy := problem
			problemCopy.Tags = nil
			if err := database.DB.Create(&problemCopy).Error; err != nil {
				return err
			}
			if err := seedProblemTags(problemCopy.ID, problem.Tags); err != nil {
				return err
			}
			status := model.UserProblemStatus{ProblemID: problemCopy.ID, UserID: "default"}
			if err := database.DB.Create(&status).Error; err != nil {
				return err
			}
		}
		log.Println("seeded mock problems")
		return nil
	}

	for _, problem := range Problems {
		if err := seedProblemTags(problem.ID, problem.Tags); err != nil {
			return err
		}
	}

	log.Println("ensured mock problem tags")
	return nil
}

func seedProblemTags(problemID string, tags []model.Tag) error {
	for _, tag := range tags {
		var persistedTag model.Tag
		if err := database.DB.Where("name = ?", tag.Name).FirstOrCreate(&persistedTag, model.Tag{Name: tag.Name}).Error; err != nil {
			return err
		}

		link := model.ProblemTag{ProblemID: problemID, TagID: persistedTag.ID}
		if err := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&link).Error; err != nil {
			return err
		}
	}
	return nil
}
