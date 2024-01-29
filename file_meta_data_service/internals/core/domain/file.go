package domain

import "time"

type File struct {
	ID        string // primary key
	UserId    string // sort/range key
	FolderId  string // sort/range key
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
