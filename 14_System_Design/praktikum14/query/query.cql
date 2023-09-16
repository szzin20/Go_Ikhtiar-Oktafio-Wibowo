-- Data Contoh (pengguna)
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    username TEXT
);

INSERT INTO users (user_id, username) VALUES (uuid(), 'tiar');
INSERT INTO users (user_id, username) VALUES (uuid(), 'tiara');
INSERT INTO users (user_id, username) VALUES (uuid(), 'tiari');

-- Ambil data dari pengguna
SELECT * FROM users;