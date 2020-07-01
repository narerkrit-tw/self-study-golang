package article

import (
	"time"
)

var fixture = []Article{
	{
		Title:    "article 1",
		Content:  "nuh uh",
		Uploaded: time.Unix(123456, 0),
	},
	{
		Title:    "article 2",
		Content:  "yup yup",
		Uploaded: time.Unix(456789, 0),
	},
}

func LoadAll() ([]Article, error) {
	return fixture, nil
}
