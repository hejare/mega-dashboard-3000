-- +goose Up
CREATE TABLE IF NOT EXISTS consultants
(
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  change_status VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS consultants CASCADE;