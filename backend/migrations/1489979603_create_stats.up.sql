CREATE TABLE stats (
  id                  SERIAL NOT NULL PRIMARY KEY,
  twitter_user_id     VARCHAR(255) NOT NULL,
  day                 DATE NOT NULL,
  impressions_count   INTEGER NOT NULL DEFAULT 0,
  new_followers_count INTEGER NOT NULL DEFAULT 0,

  UNIQUE (twitter_user_id, day)
);
