package mock

import (
	"log"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

func SeedProblems() error {
	var count int64
	if err := database.DB.Model(&model.Problem{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, problem := range Problems {
		if err := database.DB.Create(&problem).Error; err != nil {
			return err
		}
	}

	log.Println("seeded mock problems")
	return nil
}
