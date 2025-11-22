-- +goose Up
ALTER TABLE assignments ADD lead_id INTEGER REFERENCES leads(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE assignments DROP COLUMN lead_id;