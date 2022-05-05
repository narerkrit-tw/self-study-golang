package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	status int
	body   interface{}
}

func (jr JsonResponse) WriteTo(w http.ResponseWriter) {
	json, err := json.Marshal(jr.body)
	if err != nil {
		w.WriteHeader(500)
		errMessage := fmt.Sprintf("failed to encode JSON response: %s", err.Error())
		w.Write([]byte(errMessage))
	} else {
		w.WriteHeader(jr.status)
		w.Write(json)
	}
}

func NewJsonResponse(status int, body interface{}) JsonResponse {
	return JsonResponse{status: status, body: body}
}

