package main

import (
	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/service"
	"github.com/eonianmonk/url-shortener/types"
)

func main() {
	service := service.Service{
		Storage: data.NewRuntimeStorage(types.NewEncoder()),
		Port:    ":8000",
		Host:    "localhost:8000/",
	}
	service.RouteAndServe()
}
