-- +goose Up
CREATE TABLE IF NOT EXISTS assignments
(
  id SERIAL PRIMARY KEY,
  organization TEXT NOT NULL,
  contact TEXT,
  role TEXT,
  stack TEXT[],
  price NUMERIC,
  period_start_at TIMESTAMP,
  period_end_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS assignments;