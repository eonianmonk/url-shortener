package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/eonianmonk/url-shortener/data"
	"github.com/eonianmonk/url-shortener/types"
	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	cli      *redis.Client
	ctx      context.Context
	listname string
	encoder  types.Encoder
}

// urls - hashmap
// url : short url

func NewRedisStorage(addr string, password string, DB int, listname string, ctx context.Context, enc types.Encoder) data.Storage {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
	})

	if cli == nil {
		panic("failed to establish connection to redis")
	}
	// check if map exists
	t, err := cli.Type(ctx, listname).Result()
	if err != nil {
		log.Panicf("failed to check redis value type: %s", err.Error())
	}
	if t != "zset" {
		if t == "none" {
			_, err = cli.ZAdd(ctx, listname, redis.Z{Score: 0, Member: "created"}).Result()
			if err != nil {
				log.Printf("failed to create set %s for redis db: %s", listname, err.Error())
				panic("redis setup error")
			}
		} else {
			panic("redis key is occupied by other type storage")
		}
	}
	return &redisStorage{cli: cli, ctx: ctx, listname: listname, encoder: enc}
}

func (rs *redisStorage) Put(url types.Url) (types.ShortUrl, error) {
	var id int64 = 0

	err := rs.cli.Watch(rs.ctx, func(tx *redis.Tx) error {
		score, err := tx.ZScore(rs.ctx, rs.listname, string(url)).Result()
		if err != redis.Nil && err != nil {
			return fmt.Errorf("unknown error on zscore: %s", err.Error())
		}
		if err == nil { // entry exists
			id = int64(score)
			return nil
		}

		// next id = number of entries
		id, err = tx.ZCard(rs.ctx, rs.listname).Result()
		if err != nil {
			return fmt.Errorf("failed to call zcount: %s", err.Error())
		}
		_, err = tx.Pipelined(rs.ctx, func(p redis.Pipeliner) error {
			_, err = p.ZAdd(rs.ctx, rs.listname, redis.Z{
				Score: float64(id), Member: string(url),
			}).Result()
			if err != nil {
				return fmt.Errorf("failed to add entry to sorted set: %s", err.Error())
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to execute tx: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to put value to redis: %s", err.Error())
	}

	return rs.encoder.Encode(types.ID(id)), nil
}

func (rs *redisStorage) Get(surl types.ShortUrl) (types.Url, error) {
	id, err := rs.encoder.Decode(surl)
	if err != nil {
		return "", fmt.Errorf("failed to decode short url: %s", err.Error())
	}
	urlrng, err := rs.cli.ZRange(rs.ctx, rs.listname, int64(id), int64(id)).Result()
	if err != nil {
		return "", fmt.Errorf("failed to obtain zrange: %s", err.Error())
	}
	return types.Url(urlrng[0]), nil
}
