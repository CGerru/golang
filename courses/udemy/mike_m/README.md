# Build authenticated restful API with Golang

Here are the works on the Udemy course, really a quick one, just for warm up.

Packages to install:
- gorilla/mux
- dgrijalva(jwt
- /lib/pq

There's also needed a postgreSQL DB to persists users with this table

create table users (
  id serial primary key,
  email text not null unique,
  password text not null
);