CREATE TABLE keyword_suggestions (
  id            SERIAL NOT NULL PRIMARY KEY,

  campaign_id   INTEGER NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
  keyword       VARCHAR(255) NOT NULL,
  potential_impressions INTEGER NOT NULL DEFAULT 0,

  UNIQUE (campaign_id, keyword)
);
