package uploads

import (
	"context"
	"encoding/json"
	"time"
	"fmt"

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

	state.Timestamp = time.Now()

	val, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal upload state: %w", err)
	}

	return s.client.Set(ctx, key, val, 24*time.Hour).Err()
}

func (s *RedisStore) GetState(ctx context.Context, uploadID string) (UploadState, error) {
	key := "upload:" + uploadID

	val, err := s.client.Get(ctx, key).Bytes()

	if err == redis.Nil {
		return UploadState{}, fmt.Errorf("upload state not found for ID %s", uploadID)
	}
	if err != nil {
		return UploadState{}, fmt.Errorf("failed to retrieve upload state from Redis: %w", err)
	}

	var state UploadState
	if err := json.Unmarshal(val, &state); err != nil {
		return UploadState{}, fmt.Errorf("failed to unmarshal upload state: %w", err)
	}

	return state, nil
}

var _ StateStore = (*RedisStore)(nil)
