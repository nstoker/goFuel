CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz DEFAULT CURRENT_TIMESTAMP
)