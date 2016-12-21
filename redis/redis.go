package redis

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

var cli redis.Conn

func init() {
	fmt.Println("Trying to connect to redis...")
	var err error
	cli, err = redis.Dial("tcp", os.Getenv("REDIS_HOST"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to redis.")
}

func Info() {
	info, err := cli.Do("INFO")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", info)
}

func Set(table string, key string, value string) {
	_, err := cli.Do("SET", table+":"+key, value)
	if err != nil {
		panic(err)
	}
}

func Get(table string, key string) (string, error) {
	value, err := redis.String(cli.Do("GET", table+":"+key))
	return value, err
}

func Del(table string, key string) error {
	// return number of entries deleted
	_, err := cli.Do("Del", table+":"+key)
	return err
}
