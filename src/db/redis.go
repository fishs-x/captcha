package db

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"errors"
	"insur-box/src/config"
)

const (
	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

var (
	errCacheMiss = errors.New("db: key not fount")
	Redis        *RedisStore
)

type RedisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}

func init() {
	Redis = redisCache(config.RedisHost, config.RedisPassWord, config.RedisExpiration)
}

func redisCache(host string, password string, defaultExpiration time.Duration) *RedisStore {
	var pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// the redis protocol should probably be made sett-able
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	return &RedisStore{pool, defaultExpiration}
}

func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}
	conn := c.pool.Get()
	defer conn.Close()
	if expires > 0 {
		_, err := conn.Do("SETEX", key, int32(expires/time.Second), value)
		return err
	} else {
		_, err := conn.Do("SET", key, value)
		return err
	}

}

func (c *RedisStore) Get(key string) (interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	val, err := conn.Do("GET", key)
	return val, err
}

func (c *RedisStore) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return errCacheMiss
	}
	_, err := conn.Do("DEL", key)
	return err
}

func exists(conn redis.Conn, key string) bool {
	exists, _ := redis.Bool(conn.Do("EXISTS", key))
	return exists
}
