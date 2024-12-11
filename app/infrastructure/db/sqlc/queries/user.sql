-- name: FindUserByEmail :one
select id,email,name,hashed_password
from users
where email =$1;

-- name: FetchAllUser :many
select id,email,name
from users;

-- name: InsertUser :exec
insert into users (
    name,
    email
) values (
    $1,
    $2
);

-- name: UpdateUser :exec
update users
set name=$2,email=$3
where id=$1;

-- name: DeleteUser :exec
delete from users
where id=$1;