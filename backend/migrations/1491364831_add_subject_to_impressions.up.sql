ALTER TABLE impressions ADD COLUMN subject_id INTEGER;
ALTER TABLE impressions ADD COLUMN subject_type VARCHAR(255);
UPDATE impressions SET subject_id = tweet_id, subject_type = 'tweet';
ALTER TABLE impressions ALTER COLUMN subject_id SET NOT NULL;
ALTER TABLE impressions ALTER COLUMN subject_type SET NOT NULL;
ALTER TABLE impressions DROP COLUMN tweet_id;
