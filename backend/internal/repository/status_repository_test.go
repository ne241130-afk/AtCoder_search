package repository

import (
	"testing"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestStatusRepositoryCRUD(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite db: %v", err)
	}

	database.DB = db
	if err := db.AutoMigrate(&model.UserProblemStatus{}); err != nil {
		t.Fatalf("migrate status model: %v", err)
	}

	repo := &StatusRepository{}
	status, err := repo.GetStatus("abc123")
	if err != nil {
		t.Fatalf("get status: %v", err)
	}
	if status.Solved || status.Favorite || status.Memo != "" {
		t.Fatalf("expected empty default status, got %+v", status)
	}

	if err := repo.UpdateSolved("abc123", true); err != nil {
		t.Fatalf("update solved: %v", err)
	}
	if err := repo.UpdateFavorite("abc123", true); err != nil {
		t.Fatalf("update favorite: %v", err)
	}
	if err := repo.UpdateMemo("abc123", "binary search"); err != nil {
		t.Fatalf("update memo: %v", err)
	}

	status, err = repo.GetStatus("abc123")
	if err != nil {
		t.Fatalf("get status after update: %v", err)
	}
	if !status.Solved || !status.Favorite || status.Memo != "binary search" {
		t.Fatalf("expected updated status, got %+v", status)
	}
}
