sudo -i -u postgres
psql

CREATE USER groupe WITH ENCRYPTED PASSWORD 'groupe';

CREATE DATABASE groupe;

GRANT ALL PRIVILEGES ON DATABASE groupe TO groupe;

\c groupe

CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(255),
    Email VARCHAR(255),
    Username VARCHAR(255),
    Password VARCHAR(255),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

GRANT ALL PRIVILEGES ON TABLE users TO groupe;

DROP TABLE IF EXISTS group;

CREATE TABLE group (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	  REFERENCES users(id)
);

GRANT ALL PRIVILEGES ON TABLE group TO groupe;

\q
