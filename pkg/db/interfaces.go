package db

import "github.com/gomodule/redigo/redis"

type RedisClientInterface interface {
	Close() redis.Conn
	GetConnection() redis.Conn
	TestConnection() (redis.Conn, error)
}

type DBInterface interface {
	Set(prefix string, key string, value string) error
	HashSet(prefix string, key string, valueMap map[string]interface{}) error
	Get(prefix string, key string) (interface{}, error)
	GetHashSet(prefix string, key string) ([]interface{}, error)
	Exist(prefix string, key string) (bool, error)
	Delete(prefix string, key string) error
	FindHashSets(min int, max int) ([]string, error)
	SetScore(prefix string, key string, score int) error
	UpdateScore(prefix string, key string, score int) error
}
