-- +goose Up
CREATE TABLE IF NOT EXISTS leads
(
  id SERIAL PRIMARY KEY,
  organization TEXT NOT NULL,
  contact TEXT,
  role TEXT,
  stack TEXT[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS leads;