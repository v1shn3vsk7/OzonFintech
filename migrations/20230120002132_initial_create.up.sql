CREATE TABLE links (
    id bigserial not null primary key,
    origin_link varchar not null unique,
    short_link varchar not null unique
);