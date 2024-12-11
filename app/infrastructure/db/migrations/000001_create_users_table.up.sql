create table users
(
    id char(26) not null,
    email char(254) not null,
    name char(255) not null,
    password char(255) not null,
    primary key(id)
);

create index idx_email on users (email);