CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password BYTEA NOT NULL,
);

CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    path TEXT NOT NULL,
    upload_date TIMESTAMP NOT NULL,
    upload_user INT,
    category TEXT,
    FOREIGN KEY (upload_user) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    path TEXT NOT NULL,
    upload_date TIMESTAMP NOT NULL,
    upload_user INT,
    category TEXT,
    FOREIGN KEY (upload_user) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    path TEXT NOT NULL,
    upload_date TIMESTAMP NOT NULL,
    upload_user INT,
    category TEXT,
    FOREIGN KEY (upload_user) REFERENCES users(id)
);
