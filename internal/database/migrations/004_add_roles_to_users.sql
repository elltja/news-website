CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS articles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(320) NOT NULL,
    role role DEFAULT 'user',
    hashed_password TEXT NOT NULL
);

ALTER TABLE users
ADD CONSTRAINT unique_email UNIQUE (email);

CREATE TYPE IF role AS ENUM ('user', 'admin');

ALTER TABLE users
ADD COLUMN IF NOT EXISTS role role DEFAULT 'user';

UPDATE users
SET role = 'user'
WHERE role IS NULL;
