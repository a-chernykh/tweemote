CREATE MATERIALIZED VIEW followers_summary AS
  SELECT twitter_user_id, followed_at, COUNT(*) AS count
  FROM followers
  GROUP BY twitter_user_id, followed_at
  ORDER BY followed_at ASC;
