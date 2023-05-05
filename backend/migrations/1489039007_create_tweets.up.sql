CREATE TABLE tweets (
  id          SERIAL NOT NULL PRIMARY KEY,
  tweet_id    VARCHAR(255) NOT NULL,
  user_id     VARCHAR(255) NOT NULL,
  tweet_text  TEXT NOT NULL,
  created_at  TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE(tweet_id, user_id)
);
