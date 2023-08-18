package data

import (
	"fmt"
	"sync"

	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/types"
)

const (
	orderBase = 1
)

type runtimeStorage struct {
	mx      *sync.RWMutex
	encoder types.Encoder
	urlMap  map[types.Url]types.ID
	idMap   []types.Url // map[ID]Url
	order   int
}

func NewRuntimeStorage(e types.Encoder) data.Storage {
	return &runtimeStorage{
		mx:      &sync.RWMutex{},
		encoder: e,
		urlMap:  make(map[types.Url]types.ID),
		idMap:   make([]types.Url, 0),
		order:   orderBase,
	}
}

func (s *runtimeStorage) Put(url types.Url) (types.ShortUrl, error) {

	s.mx.RLock()
	id, ok := s.urlMap[url]
	s.mx.RUnlock()

	if ok {
		return s.encoder.Encode(types.ID(id)), nil
	}

	s.mx.Lock()
	id = types.ID(s.order)
	s.urlMap[url] = types.ID(s.order)
	s.idMap = append(s.idMap, url)
	s.order++
	s.mx.Unlock()

	return s.encoder.Encode(id), nil
}

func (s *runtimeStorage) Get(surl types.ShortUrl) (types.Url, error) {
	id, err := s.encoder.Decode(surl)
	if err != nil {
		return "", fmt.Errorf("failed to decode short url: %s", err.Error())
	}

	s.mx.RLock()
	defer s.mx.RUnlock()

	if s.order-orderBase < int(id) { // len(idmap) < id
		return "", fmt.Errorf("no url with id \"%s\" found", surl)
	}
	val := s.idMap[id-orderBase]

	return val, nil
}
