ALTER TABLE twitter_accounts ADD COLUMN twitter_application_id INTEGER REFERENCES twitter_applications(id) ON DELETE CASCADE;
UPDATE twitter_accounts SET twitter_application_id = (SELECT id FROM twitter_applications LIMIT 1);
ALTER TABLE twitter_accounts ALTER COLUMN twitter_application_id SET NOT NULL;
