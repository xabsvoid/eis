-- +goose Up
CREATE INDEX participants_event_id_idx ON participants(event_id);

-- +goose Down
DROP INDEX participants_event_id_idx;
