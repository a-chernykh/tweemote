CREATE TABLE campaigns (
  id                 SERIAL NOT NULL PRIMARY KEY,

  name               VARCHAR(255) NOT NULL,
  twitter_account_id INTEGER NOT NULL REFERENCES twitter_accounts(id) ON DELETE CASCADE,

  created_at         TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at         TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE (twitter_account_id, name)
);
