ALTER TABLE impressions ADD COLUMN twitter_user_id VARCHAR(255);
UPDATE impressions SET twitter_user_id = tweets.user_id FROM tweets WHERE impressions.subject_id = tweets.id AND impressions.subject_type = 'tweet';
ALTER TABLE impressions ALTER COLUMN twitter_user_id SET NOT NULL;
