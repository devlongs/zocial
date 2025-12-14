-- Enable citext extension
CREATE EXTENSION IF NOT EXISTS citext;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password BYTEA NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NOW() NOT NULL
);