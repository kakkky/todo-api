package repository

import (
	"context"
	"database/sql"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
	"github.com/kakkky/app/infrastructure/db/sqlc"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Save(ctx context.Context, user *user.User) error {
	queries := sqlc.GetQueries()
	params := sqlc.InsertUserParams{
		ID:             user.GetID(),
		Name:           user.GetName(),
		Email:          user.GetEmail().Value(),
		HashedPassword: user.GetHashedPassword().Value(),
	}
	if err := queries.InsertUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	queries := sqlc.GetQueries()
	u, err := queries.FindUserByEmail(ctx, email.Value())
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrNotFoundUser
	}
	if err != nil {
		return nil, err
	}
	user := user.ReconstructUser(
		u.ID,
		u.Email,
		u.Name,
		u.HashedPassword,
	)
	return user, nil
}

func (ur *userRepository) FetchAllUsers(ctx context.Context) (user.Users, error) {
	queries := sqlc.GetQueries()
	us, err := queries.FetchAllUser(ctx)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrNotFoundUser
	}
	if err != nil {
		return nil, err
	}
	// あらかじめ配列のキャパを確保しておく
	users := make(user.Users, 0, len(us))
	for _, u := range us {
		user := user.ReconstructUser(
			u.ID,
			u.Email,
			u.Name,
			u.HashedPassword,
		)
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepository) Update(ctx context.Context, user *user.User) error {
	queries := sqlc.GetQueries()
	params := sqlc.UpdateUserParams{
		Name:  user.GetName(),
		Email: user.GetEmail().Value(),
		ID:    user.GetID(),
	}
	if err := queries.UpdateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Delete(ctx context.Context, user *user.User) error {
	queries := sqlc.GetQueries()
	if err := queries.DeleteUser(ctx, user.GetID()); err != nil {
		return err
	}
	return nil
}
