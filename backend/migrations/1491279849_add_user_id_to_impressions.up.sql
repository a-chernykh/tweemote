ALTER TABLE impressions ADD COLUMN user_id VARCHAR(255);
UPDATE impressions SET user_id = '';
ALTER TABLE impressions ALTER COLUMN user_id SET NOT NULL;
