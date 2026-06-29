package mock

import "github.com/ne241130/atcoder-learning-hub/backend/internal/model"

var Problems = []model.Problem{
	{
		ID:         "abc042_b",
		Title:      "文字列大好き",
		Contest:    "ABC042",
		Difficulty: 140,
		Tags: []string{
			"sort",
			"string",
		},
	},
	{
		ID:         "abc153_e",
		Title:      "Crested Ibis vs Monster",
		Contest:    "ABC153",
		Difficulty: 1100,
		Tags: []string{
			"dp",
		},
	},
	{
		ID:         "abc146_c",
		Title:      "Buy an Integer",
		Contest:    "ABC146",
		Difficulty: 700,
		Tags: []string{
			"binary_search",
		},
	},
}