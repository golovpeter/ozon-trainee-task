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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS url_mappings;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
