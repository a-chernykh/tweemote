CREATE TABLE twitter_applications (
  id               SERIAL NOT NULL PRIMARY KEY,

  name             VARCHAR(255) NOT NULL,
  consumer_key     VARCHAR(255) NOT NULL,
  consumer_secret  VARCHAR(255) NOT NULL,

  created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

  UNIQUE (name),
  UNIQUE (consumer_key)
);
