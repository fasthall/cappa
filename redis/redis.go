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

func Info() {
	conn := pool.Get()
	defer conn.Close()
	info, err := conn.Do("INFO")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", info)
}

func Set(table string, key string, value string) {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", table+":"+key, value)
	if err != nil {
		panic(err)
	}
}

func Get(table string, key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", table+":"+key))
	return value, err
}

func Del(table string, key string) error {
	conn := pool.Get()
	defer conn.Close()
	// return number of entries deleted
	_, err := conn.Do("Del", table+":"+key)
	return err
}
