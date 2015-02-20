create table auth_user (
  ID SERIAL,
  Email varchar(150) not null unique,
  Password varchar(150) not null
);
