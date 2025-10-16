package uploads

import "time"

type CompletedPart struct {
	ETag       string `json:"ETag"`
	PartNumber int64  `json:"partNumber"`
}

type InitiateRequest struct {
	FileName string `json:"fileName"`
}

type PresignedURLRequest struct {
	PartNumber int64 `json:"partNumber"`
}

type CompleteRequest struct {
	Parts []CompletedPart `json:"parts"`
}

type UploadState struct {
	ID         string    `json:"id"`
	S3UploadID string    `json:"s3UploadId"`
	FileName   string    `json:"fileName"`
	Status     string    `json:"status"`
	Timestamp  time.Time `json:"timestamp"`
}