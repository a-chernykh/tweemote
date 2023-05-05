CREATE TABLE twitter_accounts (
  id                  SERIAL NOT NULL PRIMARY KEY,
  user_id             INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  twitter_user_id     VARCHAR(255) NOT NULL,
  twitter_username    VARCHAR(255) NOT NULL,
  access_token        VARCHAR(255) NOT NULL,
  access_token_secret VARCHAR(255) NOT NULL,

  UNIQUE (twitter_user_id)
);
