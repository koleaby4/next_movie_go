-- users
create table if not exists users (
    id serial primary key,
    email varchar(100) not null);

create unique index users_unique_email on users (email);

-- movies
create table if not exists movies (
    id varchar(100) primary key,
    title varchar(200) not null);

create index movies_title on movies (title);

-- user watched movies
create table if not exists movies_watched_by_user (
    user_id int not null,
    movie_id varchar(100) not null,
    experience_stars int not null,
    primary key (user_id, movie_id),
    CONSTRAINT fk_movies_watched_by_user_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_movies_watched_by_user_movie_id FOREIGN KEY(movie_id) REFERENCES movies(id),
    CONSTRAINT movies_watched_by_user_check_experience_stars CHECK(experience_stars between 1 and 5)
);
