package db

import (
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	Pool *redis.Pool
}
