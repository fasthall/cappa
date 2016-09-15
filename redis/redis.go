package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var cli redis.Conn

func init() {
	fmt.Println("Trying to connect to redis...")
	var err error
	cli, err = redis.Dial("tcp", ":6379")
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
	cli.Do("SET", table+":"+key, value)
}

func Get(table string, key string) string {
	value, _ := redis.String(cli.Do("GET", table+":"+key))
	return value
}

func Del(table string, key string) {
	cli.Do("Del", table+":"+key)
}
