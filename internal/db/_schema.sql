create table if not exists users (id serial, email varchar(100) not null);

create unique index users_unique_email on users (email);

create table if not exists movies (id varchar(100) primary key, title varchar(200) not null);

create index movies_title on movies (title);