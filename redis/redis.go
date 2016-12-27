package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool redis.Pool

func init() {
	fmt.Println("Trying to connect to redis...")
	pool = redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	fmt.Println("Connected to redis.")
}

// Info prints the information of Redis server
func Info() {
	conn := pool.Get()
	defer conn.Close()
	info, err := conn.Do("INFO")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", info)
}

// Set sets the value of the given key
func Set(table, key string, value string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("SET", table+":"+key, value)
}

// Hmset sets values of multiple fields of the given hash
func Hmset(table, key string, uuid, image, status string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("HMSET", table+":"+key, "uuid", uuid, "image", image, "status", status)
}

// Hset sets the value of the given field of hash
func Hset(table, key string, status string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("HSET", table+":"+key, "status", status)
}

// Get gets the valud of the given key
func Get(table, key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", table+":"+key))
	return value, err
}

func Hgetall(table, key string) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := conn.Do("HGETALL", table+":"+key)
	if err != nil {
		return nil, err
	}
	s, err := redis.Strings(v, err)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, err
	}
	args := map[string]string{
		s[0]: s[1],
		s[2]: s[3],
		s[4]: s[5],
	}
	return args, nil
}

// Del deletes the given key
func Del(table, key string) error {
	conn := pool.Get()
	defer conn.Close()
	// return number of entries deleted
	_, err := conn.Do("Del", table+":"+key)
	return err
}
