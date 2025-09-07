-- +goose Up
-- +goose StatementBegin
create table documents (
    id uuid primary key default gen_random_uuid(),
    filename text not null,
    mime text not null,
    is_public boolean not null,
    is_file boolean not null,
    created_at timestamptz not null default now(),
    doc_json jsonb,
    doc_url text 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table documents;
-- +goose StatementEnd
