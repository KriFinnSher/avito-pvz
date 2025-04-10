CREATE TABLE IF NOT EXISTS users (
                       id UUID PRIMARY KEY,
                       email TEXT NOT NULL UNIQUE,
                       role TEXT NOT NULL CHECK (role IN ('employee', 'moderator')),
                       hash TEXT NOT NULL
);