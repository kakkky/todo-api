create table tasks(
    id char(26) not null,
    user_id char(26) not null,
    content text not null,
    state int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    primary key(id),
    constraint fk_user_id foreign key (user_id) references users (id) on delete cascade
);

create index idx_user_id on tasks (user_id);

