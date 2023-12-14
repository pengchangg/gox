package main

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"log"
)

var (
	redisHost     string
	redisPort     int
	redisPassword string
	redisDB       int
)

func init() {
	if os.Getenv("REDIS_HOST") != "" {
		redisHost = os.Getenv("REDIS_HOST")
	} else {
		redisHost = "localhost"
	}

	if os.Getenv("REDIS_PORT") != "" {
		redisPort, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	} else {
		redisPort = 6379
	}

	if os.Getenv("REDIS_PASSWORD") != "" {
		redisPassword = os.Getenv("REDIS_PASSWORD")
	} else {
		redisPassword = ""
	}

	if os.Getenv("REDIS_DB") != "" {
		redisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	} else {
		redisDB = 0
	}

}

func main() {
	addr := redisHost + ":" + strconv.Itoa(redisPort)
	log.Printf("ridis addr : %s\n", addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	cmd := client.Conn().Ping(context.Background())

	log.Printf("Ping: %s,error:%s\n", cmd.Val(), cmd.Err().Error())
}
