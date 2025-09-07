package dto

import (
	"encoding/json"
	"os"
)

type MetaUploadRequest struct {
	Name     *string  `json:"name" validate:"required,min=1"`
	IsFile   *bool    `json:"file" validate:"required"`
	IsPublic *bool    `json:"public" validate:"required"`
	Mime     *string  `json:"mime" validate:"required,mimetype"`
	Grant    []string `json:"grant"`
}

type UploadDocumentMultipartBody struct {
	Meta MetaUploadRequest
	File *os.File
	Json json.RawMessage
}

type ErrorObj struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Envelope struct {
	Error    *ErrorObj               `json:"error,omitempty"`
	Response *map[string]interface{} `json:"response,omitempty"`
	Data     *map[string]interface{} `json:"data,omitempty"`
}

type EnvelopeOption func(*Envelope)

func NewEnvelope(opts ...EnvelopeOption) Envelope {
	s := &Envelope{}
	for _, opt := range opts {
		opt(s)
	}
	return *s
}

func WithError(code int, err error) EnvelopeOption {
	return func(s *Envelope) {
		s.Error = &ErrorObj{
			Code: code,
			Text: err.Error(),
		}
	}
}

func WithResponse(response map[string]interface{}) EnvelopeOption {
	return func(s *Envelope) {
		s.Response = &response
	}
}

func WithData(data map[string]interface{}) EnvelopeOption {
	return func(s *Envelope) {
		s.Data = &data
	}
}
