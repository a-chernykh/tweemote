CREATE TABLE keyword_stats (
  id            SERIAL NOT NULL PRIMARY KEY,

  keyword_id    INTEGER NOT NULL REFERENCES keywords(id) ON DELETE CASCADE,
  day           DATE NOT NULL,

  impressions_count   INTEGER NOT NULL DEFAULT 0,
  followers_count     INTEGER NOT NULL DEFAULT 0,

  UNIQUE (keyword_id, day)
);
