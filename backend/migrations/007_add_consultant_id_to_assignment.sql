-- +goose Up
ALTER TABLE assignment ADD consultant_id INTEGER REFERENCES consultants(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE assignment DROP COLUMN consultant_id;