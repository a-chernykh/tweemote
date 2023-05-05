CREATE TABLE users (
  id                  SERIAL NOT NULL PRIMARY KEY,
  email               VARCHAR(255) NOT NULL,
  password_hash       VARCHAR(255) NOT NULL,

  UNIQUE(email)
);
