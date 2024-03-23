-- Migration to create the "users" table
CREATE TABLE IF NOT EXISTS users(
   id SERIAL PRIMARY KEY,
   username VARCHAR(255) NOT NULL UNIQUE,
   password_hash VARCHAR(255) NOT NULL
);

-- Migration to create the "categories" table
CREATE TABLE IF NOT EXISTS categories (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   user_id INT NOT NULL,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_name_user_id ON categories (name, user_id);

CREATE INDEX IF NOT EXISTS idx_category_name ON categories (name);

-- Migration to create the "tasks" table
CREATE TABLE IF NOT EXISTS tasks (
   id SERIAL PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   description TEXT,
   status VARCHAR(255) NOT NULL DEFAULT 'todo' CHECK (status IN ('todo', 'in_progress', 'completed')),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   user_id INT NOT NULL,
   category_id INT NOT NULL,
   FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

