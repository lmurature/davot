package clients

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"translatorBot/cmd/api/utils/apierrors"
)

const (
	lastProcessedKey = "last_processed"
)

type cacheClient struct {
	redis *redis.Client
}

func (c *cacheClient) SetLastProcessedTweet(tweetID string) apierrors.ApiError {
	err := c.redis.Set(context.Background(), lastProcessedKey, tweetID, 0)
	if err != nil {
		return apierrors.NewInternalServerApiError("error while setting key value in redis db", err.Err())
	}
	return nil
}

func (c *cacheClient) GetLastProcessedTweet() (*string, apierrors.ApiError) {
	res, err := c.redis.Get(context.Background(), lastProcessedKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, apierrors.NewNotFoundApiError("No last processed tweet")
		}
		fmt.Println("Error getting last processed tweet", err)
		return nil, apierrors.NewInternalServerApiError("error getting key in redis db", err)
	}
	return &res, nil
}
