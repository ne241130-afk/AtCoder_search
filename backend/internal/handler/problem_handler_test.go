package handler

import (
	"testing"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

func TestFilterProblemsByContest(t *testing.T) {
	problems := []model.Problem{
		{Contest: "ABC042"},
		{Contest: "ARC123"},
	}

	abcResults := filterProblems(problems, "", nil, "ABC", nil, nil)
	if len(abcResults) != 1 || abcResults[0].Contest != "ABC042" {
		t.Fatalf("expected only ABC problems, got %#v", abcResults)
	}

	arcResults := filterProblems(problems, "", nil, "ARC", nil, nil)
	if len(arcResults) != 1 || arcResults[0].Contest != "ARC123" {
		t.Fatalf("expected only ARC problems, got %#v", arcResults)
	}
}

func TestFilterProblemsByDifficulty(t *testing.T) {
	problems := []model.Problem{
		{Contest: "ABC001", Difficulty: 100},
		{Contest: "ABC002", Difficulty: 200},
		{Contest: "ABC003", Difficulty: 300},
	}

	min := 150
	max := 250
	results := filterProblems(problems, "", nil, "", &min, &max)
	if len(results) != 1 || results[0].Difficulty != 200 {
		t.Fatalf("expected only the middle difficulty problem, got %#v", results)
	}
}
