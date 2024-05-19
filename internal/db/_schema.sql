-- users
create table if not exists users
(
    id    serial primary key,
    email varchar(128) not null
);

create unique index users_unique_email on users (email);

-- movies
create table if not exists movies
(
    id           int primary key,
    title        varchar(256) not null,
    release_date varchar(128) not null,
    overview     text         not null,
    poster_url   varchar(256) not null,
    trailer_url  varchar(256) not null,
    rating       float        not null,
    raw_data     text         not null,
    created_at   TIMESTAMP DEFAULT NOW()
);

create index movies_title on movies (title);

-- user watched movies
create table if not exists movies_watched_by_user
(
    user_id          int          not null,
    movie_id         varchar(128) not null,
    experience_stars int          not null,
    primary key (user_id, movie_id),
    CONSTRAINT fk_movies_watched_by_user_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_movies_watched_by_user_movie_id FOREIGN KEY (movie_id) REFERENCES movies (id),
    CONSTRAINT movies_watched_by_user_check_experience_stars CHECK (experience_stars between 1 and 5)
);

create index movies_watched_by_user_user_id on movies_watched_by_user (user_id);