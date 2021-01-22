package config

import (
	"github.com/gomodule/redigo/redis"
)

var Redis redis.Conn

func init() {
	Redis = newPool().Get()
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   15,
		IdleTimeout: 600,
		Dial: func() (redis.Conn, error) {
			/*c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}

			_, err = c.Do("AUTH", "123456")
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err*/
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
