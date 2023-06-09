CREATE TABLE IF NOT EXISTS snippets (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  content TEXT,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);