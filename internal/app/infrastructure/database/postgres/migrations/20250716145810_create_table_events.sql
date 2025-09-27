-- +goose Up
CREATE TABLE events (id BIGSERIAL PRIMARY KEY,
                     title TEXT NOT NULL,
                     date TIMESTAMPTZ NOT NULL,
                     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW());

-- +goose Down
DROP TABLE events;
