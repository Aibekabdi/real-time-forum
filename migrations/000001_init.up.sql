CREATE TABLE IF NOT EXISTS users(
    id serial not null unique, 
    nickname varchar(255) not null unique, 
    gender varchar(255) not null,
    age integer not null,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    email varchar(255) not null unique, 
    password varchar(255) not null
);