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

CREATE TABLE IF NOT EXISTS tags(
    id serial not null unique, 
    name varchar(255) not null unique
);

CREATE TABLE IF NOT EXISTS posts(
    id serial not null unique, 
    user_id integer not null,
    title varchar(255) not null,
    text varchar(255) not null,
    foreign key (user_id) references users (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS comments(
    id serial not null unique, 
    user_id integer not null,
    post_id integer not null,
    text varchar(255) not null,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (post_id) references posts (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS post_tags(
    post_id integer not null,
    tag_id integer not null,
    foreign key (post_id) references posts (id) on delete cascade,
    foreign key (tag_id) references tags (id) on delete cascade
);

CREATE TABLE posts_likes (
    id serial not null unique, 
    post_id integer not null,
    user_id integer not null,
    type integer not null,
    foreign key(post_id) references posts(id) on delete cascade,
    foreign key(user_id) references users(id) on delete cascade
);

CREATE TABLE comments_likes (
    id serial not null unique, 
    comment_id integer not null,
    user_id integer not null,
    type integer not null,
    foreign key(comment_id) references comments(id) on delete cascade,
    foreign key(user_id) references users(id) on delete cascade
);