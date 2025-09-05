package repository

import (
	"context"

	"github.com/NickVasky/docstorage/internal/models"
	"github.com/google/uuid"
)

type DocumentsRepo interface {
	Add(context.Context, models.DocumentMetadata) (uuid.UUID, error)
	GetById(context.Context, uuid.UUID) (models.DocumentMetadata, error)
	GetList(context.Context, string, string, uint64, uint64) ([]models.DocumentMetadata, error)
}
