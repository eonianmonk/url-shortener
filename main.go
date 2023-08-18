package main

import (
	data_runtime "github.com/eonianmonk/url-shortener/data/runtime"
	"github.com/eonianmonk/url-shortener/service"
	"github.com/eonianmonk/url-shortener/types"
)

func main() {
	service := service.Service{
		Storage: data_runtime.NewRuntimeStorage(types.NewEncoder()),
		Port:    ":8000",
		Host:    "localhost:8000/",
	}
	service.RouteAndServe()
}
