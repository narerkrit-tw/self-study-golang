package article

import (
	"errors"
	"time"
)

var fixture = []Article{
	{
		Title:    "article 1",
		Content:  "nuh uh",
		Uploaded: time.Unix(123456, 0),
		Category: "x",
	},
	{
		Title:    "article 2",
		Content:  "yup yup",
		Uploaded: time.Unix(456789, 0),
		Category: "y",
	},
}

func LoadAll() ([]Article, error) {
	return fixture, nil
}

func LoadCategory(category Category) ([]Article, error) {

	if category == "error" {
		return []Article{}, errors.New("boom!")
	}
	results := make([]Article, 0)

	for _, article := range fixture {
		if article.Category == category {
			results = append(results, article)
		}
	}

	return results, nil
}
