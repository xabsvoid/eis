CREATE TABLE participants (
  id SERIAL PRIMARY KEY,
  event_id INTEGER NOT NULL,
  principal_id INTEGER NOT NULL
);