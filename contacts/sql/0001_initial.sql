create table contact_contact (
  id SERIAL primary key,
  name varchar(150) not null,
  email varchar(150) not null unique
);

create table contact_group (
  id SERIAL primary key,
  name varchar(150) not null unique
);

create table contact_contactgroup (
  contact_id int references contact_contact (id),
  group_id int references contact_group (id),
  unique (contact_id, group_id)
);
