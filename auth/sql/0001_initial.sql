create table auth_user (
  id SERIAL,
  email varchar(150) not null unique,
  password varchar(150) not null
);
