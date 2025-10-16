package uploads

import "context"

type Uploader interface {
	InitiateUpload(ctx context.Context, req InitiateRequest) (UploadState, error)
}

type StateStore interface {
	SetState(ctx context.Context, state UploadState) error
	GetState(ctx context.Context, uploadID string) (UploadState, error)
}