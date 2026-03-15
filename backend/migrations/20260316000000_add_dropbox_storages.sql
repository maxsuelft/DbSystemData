-- +goose Up
-- +goose StatementBegin

CREATE TABLE dropbox_storages (
    storage_id    UUID PRIMARY KEY,
    app_key       TEXT NOT NULL,
    app_secret    TEXT NOT NULL,
    token_json    TEXT NOT NULL
);

ALTER TABLE dropbox_storages
    ADD CONSTRAINT fk_dropbox_storages_storage
    FOREIGN KEY (storage_id)
    REFERENCES storages (id)
    ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS dropbox_storages;

-- +goose StatementEnd
