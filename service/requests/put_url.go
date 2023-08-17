package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PUT host/new {"url":"google.com"}

type PutUrl struct {
	Url string `json:"url"`
}

func NewPutUrlRequest(r *http.Request) (PutUrl, error) {
	req := PutUrl{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, fmt.Errorf("failed to decode request body: %s", err.Error())
	}
	return req, nil
}
