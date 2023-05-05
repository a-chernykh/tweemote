DROP TABLE stats;

CREATE TABLE stats (
  id                  SERIAL NOT NULL PRIMARY KEY,

  campaign_id         INTEGER NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
  day                 DATE NOT NULL,
  impressions         INTEGER NOT NULL,
  followers           INTEGER NOT NULL,

  UNIQUE (campaign_id, day)
);
