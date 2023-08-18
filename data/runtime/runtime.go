package data

import (
	"fmt"

	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/types"
)

type runtimeStorage struct {
	encoder types.Encoder
	idMap   map[types.ID]types.Url
	order   int
}

func NewRuntimeStorage(e types.Encoder) data.Storage {
	return &runtimeStorage{
		encoder: e,
		idMap:   make(map[types.ID]types.Url),
		order:   1,
	}
}

func (s *runtimeStorage) Put(url types.Url) (types.ShortUrl, error) {
	s.idMap[types.ID(s.order)] = url
	surl := s.encoder.Encode(types.ID(s.order))
	s.order++
	return surl, nil
}

func (s *runtimeStorage) Get(surl types.ShortUrl) (types.Url, error) {
	id, err := s.encoder.Decode(surl)
	if err != nil {
		return "", fmt.Errorf("failed to decode short url: %s", err.Error())
	}
	val, ok := s.idMap[id]
	if !ok {
		return "", fmt.Errorf("no url with id \"%s\" found", surl)
	}
	return val, nil
}
