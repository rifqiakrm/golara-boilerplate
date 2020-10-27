package cache

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	cachePool *redis.Pool
)

func Init(pool *redis.Pool) {
	cachePool = pool
}
func PingCache() error {

	conn := cachePool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func GetCache(key string) ([]byte, error) {

	conn := cachePool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func SetCache(key string, value interface{}, ttl int) error {

	conn := cachePool.Get()
	defer conn.Close()

	val, err := json.Marshal(value)

	if err != nil {
		return fmt.Errorf("error while marshalling interface for cache")
	}

	_, err = conn.Do("SET", key, val)
	if err != nil {
		v := string(val)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	_, err = conn.Do("EXPIRE", key, ttl)

	if err != nil {
		return fmt.Errorf("error set expire key %s", key)
	}

	return err
}

func Exists(key string) (bool, error) {

	conn := cachePool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func DeleteCache(key string) error {

	conn := cachePool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func GetCacheKeys(pattern string) ([]string, error) {

	conn := cachePool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}
