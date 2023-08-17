package handlers

import (
	"net/http"

	"github.com/eonianmonk/url-shortener/service/context"
	"github.com/eonianmonk/url-shortener/service/requests"
	"github.com/eonianmonk/url-shortener/service/responses"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
	surl := requests.NewGetRequest(r)
	url, err := context.Storage(r).Get(surl)
	if err != nil {
		responses.RenderErr(w, err)
		return
	}
	http.Redirect(w, r, string(url), http.StatusMovedPermanently)
}
