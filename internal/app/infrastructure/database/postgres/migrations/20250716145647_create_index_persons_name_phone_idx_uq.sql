-- +goose Up
CREATE UNIQUE INDEX persons_name_phone_idx_uq ON persons(first_name, last_name, middle_name, phone);

-- +goose Down
DROP INDEX persons_name_phone_idx_uq;
