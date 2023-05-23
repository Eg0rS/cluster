package tokencache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	userIdKey = "userid"
	tokenKey  = "token"
	userIdTTL = 14 * 24 * time.Hour
)

type TokenCache struct {
	client *redis.Client
}

func New(addr string, password string) *TokenCache {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       1,

		ClientName: "auth-gateway",
	})
	return &TokenCache{
		client: rdb,
	}
}

func (c *TokenCache) UserDeleted(userId string) bool {
	val := c.client.Exists(context.TODO(), c.userIdWithKey(userId)).Val()
	return val == 1
}

func (c *TokenCache) AccessTokenDeleted(userId string) bool {
	val := c.client.Exists(context.TODO(), c.userIdWithKey(userId)).Val()
	return val == 1
}

func (c *TokenCache) AddWithTTL(userId string) error {
	err := c.client.Set(context.TODO(), c.userIdWithKey(userId), 1, userIdTTL).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *TokenCache) userIdWithKey(userId string) string {
	return c.withKey(userIdKey, userId)
}

func (c *TokenCache) tokenWithKey(token string) string {
	return c.withKey(tokenKey, token)
}

func (c *TokenCache) withKey(key, value string) string {
	return fmt.Sprintf("%s/%s", key, value)
}
