-- +goose Up
ALTER TABLE leads ADD consultant_id INTEGER REFERENCES consultants(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE leads DROP COLUMN consultant_id;