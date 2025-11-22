-- +goose Up
ALTER TABLE assignments ADD title TEXT;
ALTER TABLE leads ADD title TEXT;

-- +goose Down
ALTER TABLE assignments DROP COLUMN title;
ALTER TABLE leads DROP COLUMN title;