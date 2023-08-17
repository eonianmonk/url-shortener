package types

import "github.com/catinello/base62"

type Encoder interface {
	Encode(ID) ShortUrl          // returns id for a string
	Decode(ShortUrl) (ID, error) // decodes string to digit id
}

type encoder struct{}

func NewEncoder() Encoder {
	return &encoder{}
}

func (e *encoder) Encode(id ID) ShortUrl {
	return ShortUrl(base62.Encode(int(id)))
}

func (e *encoder) Decode(surl ShortUrl) (ID, error) {
	id, err := base62.Decode(string(surl))
	return ID(id), err
}
