CREATE TABLE impressions (
  id          SERIAL NOT NULL PRIMARY KEY,
  tweet_id    INTEGER NOT NULL REFERENCES tweets (id) ON DELETE CASCADE,
  action      VARCHAR(255) NOT NULL,
  created_at  TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE (tweet_id, action)
);
CREATE INDEX ON impressions (tweet_id);
