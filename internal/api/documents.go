package api

import (
	"context"
	"os"

	"github.com/NickVasky/docstorage/internal/codegen/apicodegen"
	"github.com/NickVasky/docstorage/internal/repository"
)

// Implements `DocumentsAPIServicer` interface from apicodegen package
type DocumentsAPIService struct {
	repo repository.DocumentsRepo
}

func NewDocumentsAPIService(repo repository.DocumentsRepo) apicodegen.DocumentsAPIServicer {
	s := DocumentsAPIService{repo: repo}
	return &s
}

func (s *DocumentsAPIService) ListDocuments(ctx context.Context, loginParam, keyParam, valueParam string, limitParam int32) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *DocumentsAPIService) HeadDocuments(ctx context.Context) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *DocumentsAPIService) UploadDocument(ctx context.Context, metaParam string, fileParam *os.File, jsonParam string) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}

	// TODO
	return resp, nil
}

func (s *DocumentsAPIService) GetDocument(ctx context.Context, idParam string) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *DocumentsAPIService) HeadDocument(ctx context.Context, idParam string) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *DocumentsAPIService) DeleteDocument(ctx context.Context, idParam string) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}
