-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS url_mappings
(
    id           uuid                              DEFAULT uuid_generate_v4(),
    original_url VARCHAR(2048) UNIQUE     NOT NULL,
    alias        VARCHAR(255) UNIQUE      NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS url_mappings_alias_idx ON url_mappings (alias);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP INDEX IF EXISTS url_mappings_alias_idx;
DROP TABLE IF EXISTS url_mappings;
-- +goose StatementEnd
