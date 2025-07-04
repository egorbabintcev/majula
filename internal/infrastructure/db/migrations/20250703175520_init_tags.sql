-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags (
  name TEXT,
  package_id TEXT,
  version_id TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (name, package_id, version_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags;
-- +goose StatementEnd
