create table auth_user (
  ID bigint not null primary key,
  Email char(150) not null unique,
  Password char(150) not null
);
