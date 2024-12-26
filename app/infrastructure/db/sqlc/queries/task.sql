-- name: FindTaskById :one
select id,user_id,content,state
from tasks
where id = sqlc.arg(id);

-- name: FetchTaskById :one
select t.id,u.name,t.user_id,t.content,t.state
from tasks as t inner join users as u
on t.user_id = u.id
where t.id = sqlc.arg(id);

-- name: FetchUserTasks :many
select t.id,u.name,t.user_id,t.content,t.state
from tasks as t inner join users as u
on t.user_id = u.id
where t.user_id = sqlc.arg(user_id);

-- name: FetchTaskss :many
select t.id,u.name,t.user_id,t.content,t.state
from tasks as t inner join users as u
on t.user_id = u.id;

-- name: InsertTask :exec
insert into tasks(
    id,
    user_id,
    content,
    state
)values(
    sqlc.arg(id),
    sqlc.arg(user_id),
    sqlc.arg(content),
    sqlc.arg(state)
);

-- name: UpdateTask :exec
update tasks
set state=sqlc.arg(state)
where id = sqlc.arg(id);

-- name: DeleteTask :exec
delete from tasks
where id=sqlc.arg(id);