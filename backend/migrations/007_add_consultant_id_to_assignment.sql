-- +goose Up
ALTER TABLE assignments ADD consultant_id INTEGER REFERENCES consultants(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE assignments DROP COLUMN consultant_id;