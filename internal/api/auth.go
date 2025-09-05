package api

import (
	"context"

	"github.com/NickVasky/docstorage/internal/codegen/apicodegen"
)

// Implements `AuthAPIServicer` interface from apicodegen package
type AuthAPIService struct {
}

func NewAuthAPIService() apicodegen.AuthAPIServicer {
	s := new(AuthAPIService)
	return s
}

func (s *AuthAPIService) RegisterUser(ctx context.Context, req apicodegen.RegisterRequest) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *AuthAPIService) AuthenticateUser(ctx context.Context, login string, password string) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}

func (s *AuthAPIService) LogoutUser(ctx context.Context) (apicodegen.ImplResponse, error) {
	resp := apicodegen.ImplResponse{Code: 501, Body: "Not Implemented"}
	// TODO
	return resp, nil
}
