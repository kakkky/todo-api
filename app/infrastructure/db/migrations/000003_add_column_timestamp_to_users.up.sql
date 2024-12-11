alter table users add column created_at timestamp default current_timestamp;
alter table users add column updated_at timestamp default current_timestamp on update current_timestamp;
