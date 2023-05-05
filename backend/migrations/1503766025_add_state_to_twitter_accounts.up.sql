ALTER TABLE twitter_accounts ADD COLUMN state VARCHAR(255);
UPDATE twitter_accounts SET state = 'active';
ALTER TABLE twitter_accounts ALTER COLUMN state SET NOT NULL;
