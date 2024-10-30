CREATE TABLE IF NOT EXISTS users
(
    uid VARCHAR(10) PRIMARY KEY NOT NULL,
    avatar_url VARCHAR(45) NOT NULL,
    username VARCHAR(25) NOT NULL UNIQUE,
    email VARCHAR(45) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    last_login TIMESTAMP DEFAULT NULL,
    login_attempts INTEGER DEFAULT 0,
    lock BOOLEAN DEFAULT FALSE,
    lock_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT unique_username_email UNIQUE (username, email)
) TABLESPACE pg_default;
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);