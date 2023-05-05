CREATE TABLE followers (
  id                  SERIAL NOT NULL PRIMARY KEY,
  twitter_user_id     VARCHAR(255) NOT NULL,
  twitter_follower_id VARCHAR(255) NOT NULL,
  followed_at         DATE NOT NULL,

  UNIQUE (twitter_user_id, twitter_follower_id, followed_at)
);
