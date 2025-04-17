package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zYasser/MyFitness/utils"
)

func VerifyRefreshToken(ctx context.Context, redis *redis.Client, refresh string, username string, logger *utils.Logger) bool {
	key := fmt.Sprintf("%s:%s", username, refresh)
	result := redis.Keys(ctx, key)
	return result.Err() == nil
}

func CreateRefreshToken(ctx context.Context, redis *redis.Client, token string, username string) bool {
	result := redis.SAdd(ctx, fmt.Sprintf("%s:%s", username, token), true, time.Hour*24*30*6)
	fmt.Println(result.Err())
	return true
}

func RemoveRefreshToken(ctx context.Context, redis *redis.Client, token string, username string) bool {
	result := redis.Del(ctx, fmt.Sprintf("%s:%s", username, token))
	return result.Err() == nil

}
