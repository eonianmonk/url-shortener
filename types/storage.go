package types

type Storage interface {
	Put(Url) (ShortUrl, error)
	Get(ShortUrl) (Url, error)
}
