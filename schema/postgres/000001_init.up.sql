-- Migration to create the "users" table
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   username VARCHAR(255) NOT NULL,
   password_hash VARCHAR(255) NOT NULL
);

-- Migration to create the "tasks" table
CREATE TABLE tasks (
   id SERIAL PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   description TEXT,
   status VARCHAR(255) NOT NULL DEFAULT 'todo' CHECK (status IN ('todo', 'in_progress', 'completed')),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   user_id INT,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Migration to create the "categories" table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Migration to create the "task_category" junction table
CREATE TABLE task_category (
   task_id INT,
   category_id INT,
   PRIMARY KEY (task_id, category_id),
   FOREIGN KEY (task_id) REFERENCES tasks(id),
   FOREIGN KEY (category_id) REFERENCES categories(id)
);