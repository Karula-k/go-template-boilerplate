CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  username text      NOT NULL UNIQUE,
  password  text NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);