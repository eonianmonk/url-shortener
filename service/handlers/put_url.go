package handlers

import (
	"net/http"

	"github.com/eonianmonk/url-shortener/service/context"
	"github.com/eonianmonk/url-shortener/service/requests"
	"github.com/eonianmonk/url-shortener/service/responses"
	"github.com/eonianmonk/url-shortener/types"
)

func PutUrl(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewPutUrlRequest(r)
	if err != nil {
		responses.RenderErr(w, err)
		return
	}
	surl, err := context.Storage(r).Put(types.Url(req.Url))
	if err != nil {
		responses.RenderErr(w, err)
		return
	}
	host := context.Host(r)
	resp := responses.NewPutUrlResponse(types.Url(req.Url), types.ShortUrl(surl), host)
	responses.Render(w, resp)
}
