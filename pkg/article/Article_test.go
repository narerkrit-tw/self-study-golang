package article

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestArticle_Json(t *testing.T){
	article := Article{
		Title:    "test",
		Content:  "content",
		Uploaded: time.Unix(0, 0),
	}

	result, err := json.Marshal(article)
	assert.Nil(t, err)
	jsonResult := string(result)
	assert.NotNil(t, result)
	assert.NotEmpty(t, jsonResult)
	fmt.Print(jsonResult)
}