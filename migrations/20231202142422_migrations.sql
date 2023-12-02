-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS url_mappings
(
    id            SERIAL PRIMARY KEY,
    original_url  VARCHAR(2048)            NOT NULL,
    shortened_url VARCHAR(255)             NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS url_mappings;
-- +goose StatementEnd
