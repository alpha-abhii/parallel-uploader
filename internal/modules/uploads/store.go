package uploads

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{
		client: client,
	}
}

func (s *RedisStore) SetState(ctx context.Context, state UploadState) error {
	key := "upload:" + state.ID

	val, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.client.Set(ctx, key, val, 24*time.Hour).Err()
}
