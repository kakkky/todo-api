-- name: FindUserByEmail :one
select id,email,name,hashed_password
from users
where email = sqlc.arg(email);

-- name: FindUserById :one
select id,email,name,hashed_password
from users
where id = sqlc.arg(id);

-- name: FetchAllUser :many
select id,email,name,hashed_password
from users;

-- name: InsertUser :exec
insert into users (
    id,
    name,
    email,
    hashed_password
) values (
    sqlc.arg(id),
    sqlc.arg(name),
    sqlc.arg(email),
    sqlc.arg(hashed_password)
);

-- name: UpdateUser :exec
update users
set name=sqlc.arg(name),email=sqlc.arg(email)
where id=sqlc.arg(id);

-- name: DeleteUser :exec
delete from users
where id=sqlc.arg(id);