ALTER TABLE impressions ADD COLUMN campaign_id INTEGER REFERENCES campaigns(id) ON DELETE CASCADE;
UPDATE impressions SET campaign_id = (SELECT id FROM campaigns LIMIT 1);
ALTER TABLE impressions ALTER COLUMN campaign_id SET NOT NULL;
