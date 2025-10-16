package uploads

import (
	"context"
	"time"

	"github.com/google/uuid"
)


type Service struct {
	store *RedisStore 
}

func NewService(store *RedisStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) InitiateUpload(req InitiateRequest) (string, error) {
	uploadID := uuid.New().String()

	state := UploadState{
		ID:       uploadID,
		FileName: req.FileName,
		Status:   "pending",
		Timestamp: time.Now(), 
	}

	err := s.store.SetState(context.Background(), state)
	if err != nil {
		return "", err
	}

	return uploadID, nil
}