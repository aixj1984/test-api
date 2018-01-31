package myredis

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"test-api/comm/beelog"

	"gopkg.in/redis.v5"
)

const (
	TOKENKEY = "FOXTESTING:TOKEN:"
)

type RedisCache struct {
	Cache *redis.Client
}

var redis_cache *RedisCache

func init() {
	init_RedisCache()
}

func init_RedisCache() {
	if redis_cache == nil {
		redis_cache = new(RedisCache)
	}
	cache := redis.NewClient(&redis.Options{
		Addr:     "10.0.12.104:6379",
		PoolSize: 10,
	})
	err := cache.Ping().Err()
	if err != nil {
		panic(err)
	}
	redis_cache.Cache = cache
}

func GetCache() *RedisCache {
	if redis_cache == nil {
		init_RedisCache()
	}
	return redis_cache
}

func (c *RedisCache) CheckToken(uid, token string) bool {
	key := TOKENKEY + uid
	if user, err := c.Cache.Get(key).Result(); err == nil {
		if user == token {
			return true
		}
	}
	return false
}

func (c *RedisCache) SetToken(uid string, tm int) (error, string) {
	h := md5.New()
	timestamp := time.Now().Unix()
	code := fmt.Sprintf("%v-%v", uid, timestamp)
	h.Write([]byte(code))
	token := hex.EncodeToString(h.Sum(nil))
	key := TOKENKEY + uid
	keepTime := time.Duration(tm) * time.Hour
	//存入redis
	err := c.Cache.Set(key, token, keepTime).Err()
	beelog.Debug(c.Cache.Get(key).Result())
	//判断token是否失效
	if err != nil {
		return err, ""
	}
	return nil, token
}
