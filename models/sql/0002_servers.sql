CREATE TABLE server (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  url VARCHAR(150) NOT NULL
);

CREATE TABLE server_group (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE server_servergroup (
  server_id INT REFERENCES server (id),
  group_id INT REFERENCES server_group (id),
  UNIQUE (server_id, group_id)
);
