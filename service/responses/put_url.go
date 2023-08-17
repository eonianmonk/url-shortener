package responses

import "github.com/eonianmonk/url-shortener/types"

type PutUrl struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

func NewPutUrlResponse(url types.Url, surl types.ShortUrl, host string) PutUrl {
	return PutUrl{
		Url:      string(url),
		ShortUrl: host + string(surl),
	}
}
