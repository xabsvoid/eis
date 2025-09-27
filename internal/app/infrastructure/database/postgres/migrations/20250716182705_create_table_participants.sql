-- +goose Up
CREATE TABLE participants (id BIGSERIAL PRIMARY KEY,
                           event_id BIGINT NOT NULL,
                           person_id BIGINT NOT NULL,
                           replacement_person_id BIGINT NULL,
                           check_in BOOLEAN NOT NULL DEFAULT FALSE,
                           confirmed BOOLEAN NOT NULL DEFAULT FALSE,
                           updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW());

-- +goose Down
DROP TABLE participants;
