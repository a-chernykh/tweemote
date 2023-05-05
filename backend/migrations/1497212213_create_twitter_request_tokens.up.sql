CREATE TABLE twitter_request_tokens (
  id            SERIAL NOT NULL PRIMARY KEY,

  twitter_application_id INTEGER NOT NULL REFERENCES twitter_applications(id) ON DELETE CASCADE,
  user_id                INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,

  oauth_token   VARCHAR(255) NOT NULL,
  oauth_secret  VARCHAR(255) NOT NULL,

  created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL
);
