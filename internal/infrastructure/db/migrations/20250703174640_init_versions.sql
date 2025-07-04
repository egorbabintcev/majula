-- +goose Up
-- +goose StatementBegin
CREATE TABLE versions (
  name TEXT,
  package_id TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  manifest TEXT NOT NULL CHECK(json_valid(manifest)),
  PRIMARY KEY (name, package_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE versions;
-- +goose StatementEnd
