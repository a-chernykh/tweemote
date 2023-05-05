CREATE TABLE workers (
  id          SERIAL NOT NULL PRIMARY KEY,

  name        VARCHAR(255) NOT NULL,
  hostname    VARCHAR(255) NOT NULL,
  worker_type VARCHAR(255) NOT NULL,

  started_at      TIMESTAMP NOT NULL,
  last_active_at  TIMESTAMP NOT NULL
);
