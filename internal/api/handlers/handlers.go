package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/NickVasky/docstorage/internal/api/codegen"
	"github.com/NickVasky/docstorage/internal/api/dto"
	"github.com/NickVasky/docstorage/internal/api/service"
)

const (
	maxReaderSizeMb = 512
)

// Implements `ServerInterface` interface from codegen package. Does all dirty work on requests
type ServerImpl struct {
	service service.ServiceInterface
}

func responseBuidler(w http.ResponseWriter, code int, body dto.Envelope) {
	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bodyBytes)
}

func errorBuilder(w http.ResponseWriter, code int, err error) {
	resp := dto.NewEnvelope(dto.WithError(code, err))

	bodyBytes, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bodyBytes)

}

func NewServerImpl(service service.ServiceInterface) *ServerImpl {
	s := ServerImpl{service: service}
	return &s
}

func (s *ServerImpl) ListDocuments(w http.ResponseWriter, r *http.Request, params codegen.ListDocumentsParams) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) HeadDocuments(w http.ResponseWriter, r *http.Request) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) UploadDocument(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxReaderSizeMb<<20)

	if err := r.ParseMultipartForm(32 << 20); err != nil {

		http.Error(w, "Invalid multipart form", http.StatusBadRequest)
		return
	}

	// Required field: meta
	metaStr := r.FormValue("meta")
	if metaStr == "" {
		errorBuilder(w, http.StatusBadRequest, errors.New("missing required form field 'meta'"))
		return
	}

	var metaObj dto.MetaUploadRequest
	err := json.Unmarshal([]byte(metaStr), &metaObj)
	if err != nil {
		errorBuilder(w, http.StatusBadRequest, errors.New("'meta' form field has wrong format"))
		return
	}
	err = dto.Validate.Struct(metaObj)
	if err != nil {
		errorBuilder(w, http.StatusBadRequest, err)
		return
	}

	var file *os.File
	var jsonObj json.RawMessage

	if *metaObj.IsFile {
		// handling file field errors

		f, _, err := r.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				file = nil
				errorBuilder(w, http.StatusBadRequest, errors.New("no file provided"))
				return
			} else {
				errorBuilder(w, http.StatusBadRequest, errors.New("error reading file"))
				return
			}
		} else {
			defer f.Close()

			tmp, err := os.CreateTemp("", "upload-*")
			if err != nil {
				errorBuilder(w, http.StatusBadRequest, errors.New("internal server error"))
				return
			}

			defer os.Remove(tmp.Name())
			defer tmp.Close()

			if _, err := io.Copy(tmp, f); err != nil {
				errorBuilder(w, http.StatusBadRequest, errors.New("error saving file"))
				return
			}

			file, err = os.Open(tmp.Name())
			if err != nil {
				errorBuilder(w, http.StatusBadRequest, errors.New("internal server error"))
				return
			}

			defer file.Close()
		}
	} else {
		jsonStr := r.FormValue("json")
		if len(jsonStr) == 0 {
			errorBuilder(w, http.StatusBadRequest, errors.New("no 'json' form field provided"))
			return
		}

		if err := json.Unmarshal([]byte(jsonStr), &jsonObj); err != nil {
			jsonObj = nil
			errorBuilder(w, http.StatusBadRequest, errors.New("'json' form field has wrong format"))
			return
		}

	}

	uploadBody := dto.UploadDocumentMultipartBody{
		Meta: metaObj,
		File: file,
		Json: jsonObj,
	}

	body, err := s.service.UploadDocument(r.Context(), uploadBody)
	if err != nil {
		errorBuilder(w, http.StatusBadRequest, errors.New("internal server error"))
		return
	}

	responseBuidler(w, http.StatusCreated, body)
}

func (s *ServerImpl) GetDocument(w http.ResponseWriter, r *http.Request, id string) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) HeadDocument(w http.ResponseWriter, r *http.Request, id string) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) DeleteDocument(w http.ResponseWriter, r *http.Request, id string) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}

func (s *ServerImpl) LogoutUser(w http.ResponseWriter, r *http.Request) {
	errorBuilder(w, http.StatusNotImplemented, errors.New("not implemented"))
	// TODO
}
