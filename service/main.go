package service

import (
	"net/http"

	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/service/context"
	"github.com/eonianmonk/url-shortener/service/handlers"
)

type Service struct {
	Storage data.Storage
	Host    string
	Port    string
}

func (s *Service) RouteAndServe() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetUrl)
	mux.HandleFunc("/new", handlers.PutUrl)

	r := CompileMiddleware(
		mux,
		context.CtxStorage(s.Storage),
		context.CtxHost(s.Host),
	)

	http.ListenAndServe(s.Port, r)
}
