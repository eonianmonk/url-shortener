package requests

import (
	"net/http"

	"github.com/eonianmonk/url-shortener/types"
)

// host/{short url}

func NewGetRequest(r *http.Request) types.ShortUrl {
	//reqPath := strings.Split(r.URL.)
	surl := r.URL.Path[1:] // skipping /

	return types.ShortUrl(surl)
}
