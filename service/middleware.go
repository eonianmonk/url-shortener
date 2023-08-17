package service

import (
	"net/http"

	"github.com/eonianmonk/url-shortener/service/context"
)

func CompileMiddleware(h http.Handler, mws ...context.Middleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		for _, mw := range mws {
			ctx = mw(ctx)
		}
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
