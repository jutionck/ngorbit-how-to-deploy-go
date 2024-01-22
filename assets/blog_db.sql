CREATE TABLE blogs (
  id SERIAL PRIMARY KEY,
  user_id int,
  title VARCHAR(100),
  body TEXT,
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);