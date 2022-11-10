CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    email varchar(255) not null unique,
    username varchar(255) not null unique,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    password_hash varchar(255) not null,
    age varchar(255) not null,
    gender varchar(255) not null
);

CREATE TABLE posts
(
    id SERIAL PRIMARY KEY,
    creator_id int not null,
    title  varchar(255) not null,
    tags  varchar(255) not null,
    content  varchar(255) not null,
    FOREIGN KEY (creator_id) REFERENCES users ON DELETE CASCADE
);

CREATE TABLE sessions
(
    id SERIAL PRIMARY KEY,
    uuid varchar(255) not null,
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE
);