package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/xpoh/BRE-test/cmd/go"
)

type CacheSorage struct {
	ctx context.Context
	rdb *redis.Client
}

func (c *CacheSorage) AddToStorage(content swagger.Content) error {
	err := c.rdb.Set(c.ctx, content.Id, content.Body, 0).Err()
	return err
}

func (c *CacheSorage) ReadFromStorage(id string) (string, error) {
	val, err := c.rdb.Get(c.ctx, id).Result()
	return val, err
}

func (c *CacheSorage) RemoveFromStorage(id string) error {
	err := c.rdb.Del(c.ctx, id).Err()
	return err
}

func NewCacheSorage(ctx context.Context) *CacheSorage {
	c := &CacheSorage{ctx: ctx}
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return c
}
