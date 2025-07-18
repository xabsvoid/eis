CREATE TABLE principal (
  id SERIAL PRIMARY KEY,
  personal_data_id INTEGER NOT NULL,
  -- no FK as requested
  UNIQUE(personal_data_id)
);