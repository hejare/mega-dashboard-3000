-- +goose Up
ALTER TABLE assignments ADD title;
ALTER TABLE leads ADD title;

-- +goose Down
ALTER TABLE assignments DROP COLUMN title;
ALTER TABLE leads DROP COLUMN title;