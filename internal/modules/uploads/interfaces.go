package uploads

import "context"

type Uploader interface {
	InitiateUpload(ctx context.Context, req InitiateRequest) (UploadState, error)
	GetPresignedURL(ctx context.Context, uploadID string, partNumber int64) (string, error)
	CompleteUpload(ctx context.Context, uploadID string, parts []CompletedPart) error
}

type StateStore interface {
	SetState(ctx context.Context, state UploadState) error
	GetState(ctx context.Context, uploadID string) (UploadState, error)
}