package uploads

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Uploader struct {
	s3Client *s3.Client
	store    StateStore
	bucket   string
}

func NewS3Uploader(s3Client *s3.Client, store StateStore, bucketName string) *S3Uploader {
	return &S3Uploader{
		s3Client: s3Client,
		store:    store,
		bucket:   bucketName,
	}
}

func (u *S3Uploader) InitiateUpload(ctx context.Context, req InitiateRequest) (UploadState, error) {
	uploadID := uuid.New().String()

	s3Input := &s3.CreateMultipartUploadInput{
		Bucket: &u.bucket,
		Key:    &req.FileName,
	}
	log.Printf("Initiating multipart upload for file: %s in bucket: %s", req.FileName, u.bucket)

	s3Output, err := u.s3Client.CreateMultipartUpload(ctx, s3Input)
	if err != nil {
		return UploadState{}, fmt.Errorf("s3 client failed to create multipart upload: %w", err)
	}

	state := UploadState{
		ID:         uploadID,
		S3UploadID: *s3Output.UploadId,
		FileName:   req.FileName,
		Status:     "pending",
		Timestamp:  time.Now(),
		Parts:      make(map[int32]string),
	}

	if err := u.store.SetState(ctx, state); err != nil {
		return UploadState{}, fmt.Errorf("failed to save initial upload state (S3 Upload ID: %s): %w", *s3Output.UploadId, err)
	}

	log.Printf("Successfully initiated S3 upload: %s (Internal ID: %s)", *s3Output.UploadId, uploadID)
	return state, nil
}

func (u *S3Uploader) GetPresignedURL(ctx context.Context, uploadID string, partNumber int64) (string, error) {
	state, err := u.store.GetState(ctx, uploadID)
	if err != nil {
		return "", fmt.Errorf("failed to get upload state for presigned URL: %w", err)
	}

	pNum := int32(partNumber)
	presignClient := s3.NewPresignClient(u.s3Client)

	uploadPartInput := &s3.UploadPartInput{
		Bucket:     &u.bucket,
		Key:        &state.FileName,
		UploadId:   &state.S3UploadID,
		PartNumber: &pNum,
	}

	presignedURL, err := presignClient.PresignUploadPart(ctx, uploadPartInput, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	log.Printf("Generated presigned URL for Part %d of Upload ID %s", partNumber, uploadID)
	return presignedURL.URL, nil
}

var _ Uploader = (*S3Uploader)(nil)
