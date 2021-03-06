package storages

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/vinhha96/golang-research/models"
	"log"
	"sync"
	"time"
)

var once sync.Once
var redisClient *RedisClient

type RedisClient struct {
	*redis.Client
}

func GetRedisClient(address, password string) *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password, // no password set
			DB:       0,        // use default DB
		})

		redisClient = &RedisClient{client}
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("[Error] Could not connect to redis %v", err)
	}

	return redisClient
}

func (c *RedisClient) SaveToStore() {
	value, errs := json.Marshal(models.User{Username: "ABC"})
	if errs != nil {
		log.Println("Marshal error")
	}

	_, err := c.Set("ABC", value, time.Hour).Result()
	if err != nil {
		log.Println("Set value error")
	}

	result, _ := c.Get("ABC").Result()

	var user models.User
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		log.Println("Get value error")
	}

	log.Println(user)
}
