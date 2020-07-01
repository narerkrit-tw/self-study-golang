package api

import (
	"encoding/json"
	"net/http"
)

type HomepageHandler struct {
}

func (h HomepageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("home page body")
}
