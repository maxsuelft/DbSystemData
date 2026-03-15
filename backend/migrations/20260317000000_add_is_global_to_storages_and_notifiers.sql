-- +goose Up
-- +goose StatementBegin
ALTER TABLE storages
    ADD COLUMN is_global BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE notifiers
    ADD COLUMN is_global BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE storages
    DROP COLUMN is_global;
ALTER TABLE notifiers
    DROP COLUMN is_global;
-- +goose StatementEnd
