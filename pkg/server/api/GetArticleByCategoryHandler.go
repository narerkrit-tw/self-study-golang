package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"self-study-golang/pkg/article"
)

func GetArticleByCategoryHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	cat := article.Category(params["category"])
	articles, err := article.LoadCategory(cat)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(articles)
}