package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	appRedisConfig "github.com/jiaqi-yin/go-file-processor/config"
	"github.com/jiaqi-yin/go-file-processor/domain"
)

type redisService struct {
	Client *redis.Client
}

func (rs *redisService) Save(item *domain.Item) {
	ctx := context.Background()
	err := rs.Client.HSet(
		ctx,
		fmt.Sprintf("users:%s", item.UUID),
		[]string{"firstname", item.Firstname, "lastname", item.Lastname},
	).Err()
	if err != nil {
		panic(err)
	}
}

func NewRedisService() Storage {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", appRedisConfig.Configurations.Redis.Host, appRedisConfig.Configurations.Redis.Port),
		Password: appRedisConfig.Configurations.Redis.Password,
		DB:       appRedisConfig.Configurations.Redis.Db,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return &redisService{
		Client: client,
	}
}
