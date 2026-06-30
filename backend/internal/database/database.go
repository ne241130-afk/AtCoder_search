package database

import (
	"fmt"
	"log"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"postgres",
		"postgres",
		"postgres",
		"atcoder_search",
		"5432",
	)

	log.Printf("connecting to postgres: host=%s dbname=%s port=%s", "postgres", "atcoder_search", "5432")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to open postgres connection: %v", err)
		return err
	}

	log.Printf("running auto-migration for models %T and %T", &model.Problem{}, &model.UserProblemStatus{})
	if err := DB.AutoMigrate(&model.Problem{}, &model.UserProblemStatus{}); err != nil {
		log.Printf("failed to auto-migrate problem schema: %v", err)
		return err
	}

	log.Printf("database connected and schema initialized")
	return nil
}
