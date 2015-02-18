create table auth_user (
  ID SERIAL,
  Email char(150) not null unique,
  Password char(150) not null
);
