package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"self-study-golang/pkg/article"
)

/*
	uses
	url var: category
 */
func GetArticleByCategoryHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	cat := article.Category(params["category"])
	var res JsonResponse
	articles, err := article.LoadCategory(cat)
	if err != nil {
		res = NewJsonResponse(500, err.Error())
	} else {
		res = NewJsonResponse(200, articles )
	}
	res.WriteTo(w)
}