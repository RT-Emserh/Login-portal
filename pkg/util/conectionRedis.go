package util

import (
	"github.com/gomodule/redigo/redis"
)

var RedisPool = &redis.Pool{
	MaxActive: 30,
	MaxIdle:   15,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.DialURL("redis://localhost:6379")
	},
}
