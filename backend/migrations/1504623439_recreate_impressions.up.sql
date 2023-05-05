DROP TABLE impressions;

CREATE TABLE impressions (
  id            SERIAL NOT NULL PRIMARY KEY,

  actor_twitter_user_id   VARCHAR(255) NOT NULL,
  subject_twitter_user_id VARCHAR(255) NOT NULL,

  action        VARCHAR(255) NOT NULL,
  subject_id    VARCHAR(255) NOT NULL,
  subject_type  VARCHAR(255) NOT NULL,

  created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL,

  UNIQUE (subject_id, subject_type, action)
);
