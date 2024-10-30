CREATE TABLE IF NOT EXISTS users
(
    uid TEXT PRIMARY KEY NOT NULL,
    avatar_url TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    last_login DATETIME DEFAULT NULL,
    login_attempts INTEGER DEFAULT 0,
    lock BOOLEAN DEFAULT FALSE,
    lock_at DATETIME DEFAULT NULL,
    UNIQUE (username, email)
);
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);