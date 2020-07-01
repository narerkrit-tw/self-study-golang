package article

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestArticle_Json(t *testing.T){
	article := Article{
		Title:    "test",
		Content:  "content",
		Uploaded: time.Unix(0, 0),
		Category: "test_category",
	}

	result, err := json.Marshal(article)
	assert.Nil(t, err)
	jsonResult := string(result)
	assert.NotNil(t, result)
	assert.NotEmpty(t, jsonResult)
	json, _ := simplejson.NewJson([]byte(jsonResult))
	titleString, _ := json.GetPath("title").String()
	assert.Equal(t, titleString, "test")
}