CREATE TABLE keyword_tweets (
  id            SERIAL NOT NULL PRIMARY KEY,

  keyword_id    INTEGER NOT NULL REFERENCES keywords(id) ON DELETE CASCADE,
  tweet_id      INTEGER NOT NULL REFERENCES tweets(id) ON DELETE CASCADE,

  created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE (keyword_id, tweet_id)
);
