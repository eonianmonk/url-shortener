package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, val interface{}) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(val)
	if err != nil {
		panic(fmt.Sprintf("failed to render response: %s", err.Error()))
	}
}

func RenderErr(w http.ResponseWriter, err error) {
	body := newErrorResponseE(err)
	Render(w, body)
}
