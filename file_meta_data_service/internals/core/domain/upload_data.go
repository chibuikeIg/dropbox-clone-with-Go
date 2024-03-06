package domain

import "time"

type UploadData struct {
	UserId    string
	UploadId  string
	Etag      string
	PartNum   int32
	CreatedAt time.Time
}
