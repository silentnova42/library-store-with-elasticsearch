CREATE TABLE IF NOT EXISTS user_profile (
    id        SERIAL       PRIMARY KEY,
    email     VARCHAR(255) NOT NULL,
    username  VARCHAR(255) NOT NULL
);