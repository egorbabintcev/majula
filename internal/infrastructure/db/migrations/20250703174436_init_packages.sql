-- +goose Up
-- +goose StatementBegin
CREATE TABLE packages (
  name TEXT PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE packages;
-- +goose StatementEnd
