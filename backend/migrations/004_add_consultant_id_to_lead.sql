-- +goose Up
ALTER TABLE leads ADD consultant_id INTEGER REFERENCES consultants(id) ON DELETE SET NULL;
