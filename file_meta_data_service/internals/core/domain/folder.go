package domain

import "time"

type Folder struct {
	ID        string // primary key
	UserId    string // sort/range key
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
