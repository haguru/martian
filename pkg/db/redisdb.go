package db

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	client RedisClientInterface
}

func NewRedisDB(client RedisClientInterface) (DBInterface, error) {
	db := &RedisDB{
		client: client,
	}

	connection, err := db.client.TestConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = connection.Close()
	}()

	return db, nil
}

func (db *RedisDB) Set(prefix string, key string, value string) error {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return errors.New("no prefix given")
	}

	_, err := connection.Do("SET", prefix+key, value)

	return err
}

func (db *RedisDB) HashSet(prefix string, key string, valueMap map[string]interface{}) error {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return errors.New("no prefix given")
	}

	_ = connection.Send("MULTI")
	_ = connection.Send("HSET", redis.Args{prefix + key}.AddFlat(valueMap)...)
	_, err := connection.Do("EXEC")

	return err

}

func (db *RedisDB) Get(prefix string, key string) (interface{}, error) {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return nil, errors.New("no prefix given")
	}

	reply, err := redis.String(connection.Do("GET", prefix+key))
	if err != nil {
		return nil, err
	}

	return reply, nil

}

func (db *RedisDB) GetHashSet(prefix string, key string) ([]interface{}, error) {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return nil, errors.New("no prefix given")
	}

	reply, err := redis.Values(connection.Do("HGETALL", prefix+key))
	if err != nil {
		return nil, err
	}

	return reply, nil

}

func (db *RedisDB) Exist(prefix string, key string) (bool, error) {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return false, errors.New("no prefix given")
	}

	reply, err := redis.Bool(connection.Do("EXISTS", prefix+key))
	if err != nil {
		return false, err
	}

	return reply, nil
}

func (db *RedisDB) Delete(prefix string, key string) error {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return errors.New("no prefix given")
	}

	_, err := connection.Do("DEL", prefix+key)

	return err
}

func (db *RedisDB) FindHashSets(min int, max int) ([]string, error) {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	reply, err := redis.Strings(connection.Do("ZRANGEBYSCORE", "score", min, max))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (db *RedisDB) SetScore(prefix string, key string, score int) error {
	connection := db.client.GetConnection()
	defer func() {
		_ = connection.Close()
	}()

	if prefix == "" {
		return errors.New("no prefix given")
	}

	_, err := connection.Do("ZADD", "score", prefix+key)

	return err
}
func (db *RedisDB) UpdateScore(prefix string, key string, score int) error {
	return db.SetScore(prefix, key, score)
}
