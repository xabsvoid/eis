CREATE TABLE representative (
  id SERIAL PRIMARY KEY,
  personal_data_id INTEGER NOT NULL,
  UNIQUE(personal_data_id)
);