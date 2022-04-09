package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	once sync.Once
)

type Configuration struct {
	Host      string
	Port      int
	BatchSize int
	Timeout   time.Duration
	Password  string //consider using vault of osme type of secure way to get it maybe create an interface for this ATM
}

type RedisClient struct {
	Pool      *redis.Pool
	BatchSize int
}

func NewClient(config Configuration) (*RedisClient, error) {
	var currClient *RedisClient
	once.Do(func() {
		connectionString := fmt.Sprintf("%s:%d", config.Host, config.Port)
		opts := []redis.DialOption{
			redis.DialConnectTimeout(time.Duration(config.Timeout) * time.Millisecond),
		}

		dialFunc := func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp", connectionString, opts...,
			)
			if err == nil {
				_, err = conn.Do("PING")
				if err == nil {
					return conn, nil
				}
			}

			return nil, fmt.Errorf("could not dial Redis: %s", err)
		}
		// Default the batch size to 1,000 if not set
		batchSize := 1000
		if config.BatchSize != 0 {
			batchSize = config.BatchSize
		}
		currClient = &RedisClient{
			Pool: &redis.Pool{
				IdleTimeout: 0,
				MaxIdle:     10,
				Dial:        dialFunc,
			},
			BatchSize: batchSize,
			//loggingClient: lc,
		}
	})

	// Test connectivity now so don't have failures later when doing lazy connect.
	if _, err := currClient.Pool.Dial(); err != nil {
		return nil, err
	}

	return currClient, nil
}

func (client *RedisClient) GetConnection() redis.Conn {
	return client.Pool.Get()
}

func (client *RedisClient) TestConnection() (redis.Conn, error) {
	connection, err := client.Pool.Dial()
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (client *RedisClient) Close() error {
	return client.Pool.Close()
}
