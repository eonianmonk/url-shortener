package main

import (
	"context"

	data_redis "github.com/eonianmonk/url-shortener/data/redis"
	//data_runtime "github.com/eonianmonk/url-shortener/data/runtime"
	"github.com/eonianmonk/url-shortener/service"
	"github.com/eonianmonk/url-shortener/types"
)

func main() {
	encoder := types.NewEncoder()
	redis := data_redis.NewRedisStorage("localhost:6379", "", 0, "urls", context.Background(), encoder)
	//runtime :=  data_runtime.NewRuntimeStorage(types.NewEncoder()),

	service := service.Service{
		Storage: redis,
		Port:    ":8000",
		Host:    "localhost:8000/",
	}
	service.RouteAndServe()
}
