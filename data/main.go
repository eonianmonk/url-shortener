package data

import "github.com/eonianmonk/url-shortener/types"

type Storage interface {
	Put(types.Url) (types.ShortUrl, error)
	Get(types.ShortUrl) (types.Url, error)
}
