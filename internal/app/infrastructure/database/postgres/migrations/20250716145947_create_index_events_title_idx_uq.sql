-- +goose Up
CREATE UNIQUE INDEX events_title_idx_uq ON events(title);

-- +goose Down
DROP INDEX events_title_idx_uq;
