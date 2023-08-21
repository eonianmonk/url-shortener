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
	url := types.Url(req.Url)
	url.Verify()
	surl, err := context.Storage(r).Put(url)
	if err != nil {
		responses.RenderErr(w, err)
		return
	}
	host := context.Host(r)
	resp := responses.NewPutUrlResponse(types.Url(url), types.ShortUrl(surl), host)
	responses.Render(w, resp)
}
