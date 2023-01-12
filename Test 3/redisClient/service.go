package redisClient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

func ConnectRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func SetData(rds *redis.Client, key string, data string) bool {

	ttl := time.Duration(10) * time.Minute

	op1 := rds.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		// fmt.Printf("unable to SET data. error: %v", err)
		return false
	}
	// log.Println("set operation success")

	return true

}

func GetData(rds *redis.Client, key string) string {

	// get data
	op2 := rds.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		// fmt.Printf("unable to GET data. error: %v", err)
		return "false"
	}
	res, err := op2.Result()
	if err != nil {
		// fmt.Printf("unable to GET data. error: %v", err)
		return "false"
	}
	// log.Println("get operation success. result:", res)

	return res

}
