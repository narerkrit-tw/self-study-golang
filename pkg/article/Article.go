package article

import "time"

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Uploaded time.Time `json:"uploaded"`
	Category Category `json:"category"`
}


type Category string