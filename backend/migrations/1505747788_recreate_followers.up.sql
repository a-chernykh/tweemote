DROP MATERIALIZED VIEW IF EXISTS followers_summary;
DROP TABLE followers;

CREATE TABLE followers (
  id                  SERIAL NOT NULL PRIMARY KEY,
  campaign_id         INTEGER REFERENCES campaigns(id) ON DELETE SET NULL,
  twitter_user_id     VARCHAR(255) NOT NULL,
  twitter_follower_id VARCHAR(255) NOT NULL,
  followed_at         DATE NOT NULL,

  UNIQUE (twitter_user_id, twitter_follower_id)
);
