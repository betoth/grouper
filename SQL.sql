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

\q
