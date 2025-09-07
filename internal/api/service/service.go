package service

import (
	"context"

	"github.com/NickVasky/docstorage/internal/api/dto"
	"github.com/NickVasky/docstorage/internal/repository"
)

type ServiceInterface interface {
	UploadDocument(ctx context.Context, body dto.UploadDocumentMultipartBody) (dto.Envelope, error)
}

type ServiceImpl struct {
	repo repository.DocumentsRepo
}

// Implements Service interface. Contains business logic
func NewServiceImpl(repo repository.DocumentsRepo) *ServiceImpl {
	s := ServiceImpl{
		repo: repo,
	}
	return &s
}

func (s *ServiceImpl) UploadDocument(ctx context.Context, body dto.UploadDocumentMultipartBody) (dto.Envelope, error) {
	resp := dto.Envelope{}
	//s.repo.Add(ctx)

	if *body.Meta.IsFile {
		resp.Data = &map[string]interface{}{
			"file": body.Meta.Name,
		}
	} else {
		resp.Data = &map[string]interface{}{
			"json": body.Json,
		}
	}
	return resp, nil
}
