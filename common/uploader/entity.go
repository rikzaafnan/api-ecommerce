package uploader

import "time"

type Attachment struct {
	ID            int
	FileName      string
	FileExtension string
	FileLocation  string
	Module        string
	UserID        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
