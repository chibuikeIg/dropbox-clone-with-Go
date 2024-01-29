package domain

import "time"

type UploadPart struct {
	UserId    string
	UploadId  string
	Etag      string
	PartNum   int32
	CreatedAt time.Time
}
