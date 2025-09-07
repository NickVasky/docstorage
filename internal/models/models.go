package models

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type DocumentMetadata struct {
	ID        uuid.UUID
	Name      string
	IsFile    bool
	MimeType  string
	IsPublic  bool
	CreatedAt time.Time
	URL       string
	JsonDoc   []byte
}

type Document struct {
	DocumentMetadata
	File os.File
}
