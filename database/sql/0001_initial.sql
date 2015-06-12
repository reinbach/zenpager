CREATE TABLE migrate (
   id SERIAL PRIMARY KEY,
   app varchar(50) not null,
   name varchar(150) not null,
   datetime timestamp not null
);
