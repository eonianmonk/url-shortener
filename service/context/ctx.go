package context

import (
	"context"
	"net/http"

	"github.com/eonianmonk/url-shortener/types"
)

type ctxKey int

type Middleware func(context.Context) context.Context

const (
	StorageCtxKey ctxKey = iota
	LinkCtxKey
)

func CtxStorage(storage types.Storage) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, StorageCtxKey, storage)
	}
}

func Storage(r *http.Request) types.Storage {
	return r.Context().Value(StorageCtxKey).(types.Storage)
}

func CtxHost(link string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, LinkCtxKey, link)
	}
}

func Host(r *http.Request) string {
	return r.Context().Value(LinkCtxKey).(string)
}
