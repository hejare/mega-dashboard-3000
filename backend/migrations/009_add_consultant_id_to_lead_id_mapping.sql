-- +goose Up
ALTER TABLE leads DROP COLUMN IF EXISTS consultant_id;

CREATE TABLE IF NOT EXISTS consultant_id_to_lead_id_mapping
(
  consultant_id INT NOT NULL,
  lead_id INT NOT NULL,
  PRIMARY KEY (consultant_id, lead_id),
  FOREIGN KEY (consultant_id) REFERENCES consultants(id) ON DELETE CASCADE,
  FOREIGN KEY (lead_id) REFERENCES leads(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS consultant_id_to_lead_id_mapping;