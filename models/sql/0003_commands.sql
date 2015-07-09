CREATE TABLE command (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  command VARCHAR(150) NOT NULL
);

CREATE TABLE command_group (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE command_commandgroup (
  command_id INT REFERENCES command (id),
  group_id INT REFERENCES command_group (id),
  UNIQUE (command_id, group_id)
);
