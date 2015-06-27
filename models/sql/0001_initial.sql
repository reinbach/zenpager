CREATE TABLE auth_user (
  id SERIAL PRIMARY KEY,
  email VARCHAR(150) NOT NULL UNIQUE,
  password VARCHAR(150) NOT NULL
);

CREATE TABLE auth_token (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(150) NOT NULL,
  token VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE contact_contact (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  user_id INT REFERENCES auth_user (id),
  UNIQUE (user_id)
);

CREATE TABLE contact_group (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE contact_contactgroup (
  contact_id INT REFERENCES contact_contact (id),
  group_id INT REFERENCES contact_group (id),
  UNIQUE (contact_id, group_id)
);
