-- +goose Up
CREATE TABLE persons (id BIGSERIAL PRIMARY KEY,
                      first_name TEXT NOT NULL,
                      last_name TEXT NOT NULL,
                      middle_name TEXT NOT NULL,
                      phone TEXT NOT NULL,
                      updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW());

-- +goose Down
DROP TABLE persons;
