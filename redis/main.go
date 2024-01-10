package main

import (
	"context"
	"os"
	"strconv"

	"log"

	"github.com/redis/go-redis/v9"
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

var ctx = context.Background()
var key = "key"

func main() {
	addr := redisHost + ":" + strconv.Itoa(redisPort)
	log.Printf("ridis addr : %s,password:%s,DB:%d\n", addr, redisPassword, redisDB)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	err := rdb.Ping(ctx).Err()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("redis connect success")

	log.Println(
		rdb.Set(ctx, key, "value", -1).Err(),
		rdb.Get(ctx, key).Val(),
	)

}
