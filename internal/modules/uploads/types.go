package uploads

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