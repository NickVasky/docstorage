package documents

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/NickVasky/docstorage/internal/models"
	"github.com/NickVasky/docstorage/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "documents"

	idColumn        = "id"
	filenameColumn  = "filename"
	mimeColumn      = "mime"
	isPublicColumn  = "is_public"
	isFileColumn    = "is_file"
	createdAtColumn = "created_at"
	urlColumn       = "doc_url"
	docJsonColumn   = "doc_json"
)

var psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

var ErrWrongKey = fmt.Errorf("wrong filter key")

var allowedFilterKeys = map[string]struct{}{
	mimeColumn:     {},
	isFileColumn:   {},
	isPublicColumn: {},
}

type repo struct {
	dbc *pgxpool.Pool
}

func NewRepo(dbc *pgxpool.Pool) repository.DocumentsRepo {
	return &repo{dbc: dbc}
}

func (r *repo) Add(ctx context.Context, meta models.DocumentMetadata) (uuid.UUID, error) {
	var id uuid.UUID
	builder := psq.
		Insert(tableName).
		Columns(idColumn, filenameColumn, mimeColumn, isPublicColumn, isFileColumn, createdAtColumn, urlColumn, docJsonColumn).
		Values(meta.ID, meta.Name, meta.MimeType, meta.IsPublic, meta.IsFile, meta.CreatedAt, meta.URL, meta.JsonDoc).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return id, err
	}

	err = r.dbc.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *repo) GetById(ctx context.Context, id uuid.UUID) (models.DocumentMetadata, error) {
	var doc models.DocumentMetadata

	builder := psq.
		Select(idColumn, filenameColumn, mimeColumn, isPublicColumn, isFileColumn, createdAtColumn, urlColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return doc, err
	}

	err = r.dbc.QueryRow(ctx, query, args...).Scan(&doc)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (r *repo) GetList(ctx context.Context, key string, value string, limit uint64, offset uint64) ([]models.DocumentMetadata, error) {
	docs := make([]models.DocumentMetadata, 0, limit)

	if _, ok := allowedFilterKeys[key]; !ok {
		return docs, ErrWrongKey
	}

	builder := psq.
		Select(idColumn, filenameColumn, mimeColumn, isPublicColumn, isFileColumn, createdAtColumn, urlColumn).
		From(tableName).
		Where(sq.Eq{key: value}).Limit(uint64(limit)).Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return docs, err
	}

	rows, err := r.dbc.Query(ctx, query, args...)
	if err != nil {
		return docs, err
	}

	docs, err = pgx.CollectRows(rows, pgx.RowTo[models.DocumentMetadata])
	if err != nil {
		return docs, err
	}

	return docs, nil
}
