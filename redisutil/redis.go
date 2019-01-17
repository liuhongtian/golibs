package redisutil

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetFromRedis(key string) string {
	conn, err := redis.Dial("tcp", "redis.default:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return ""
	}
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got liu %v \n", value)
	}

	return value
}
