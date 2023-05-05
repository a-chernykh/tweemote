CREATE TABLE keywords (
  id                 SERIAL NOT NULL PRIMARY KEY,

  keyword            VARCHAR(255) NOT NULL,
  campaign_id        INTEGER NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,

  created_at         TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at         TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE (campaign_id, keyword)
);
