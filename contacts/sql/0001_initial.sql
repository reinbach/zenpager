CREATE TABLE contact_contact (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  user_id INT REFERENCES auth_user (id)
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
