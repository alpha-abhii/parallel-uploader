package uploads

import "github.com/google/uuid"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) InitiateUpload(req InitiateRequest) (string, error) {
	uploadID := uuid.New().String()

	return uploadID, nil
}