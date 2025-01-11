package repository

import (
	"context"
	"database/sql"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/user"
)

type userRepository struct {
	querier Querier
}

func NewUserRepository(querier Querier) user.UserRepository {
	return &userRepository{
		querier: querier,
	}
}

func (ur *userRepository) Save(ctx context.Context, user *user.User) error {
	arg := InsertUserParams{
		ID:             user.GetID(),
		Name:           user.GetName(),
		Email:          user.GetEmail().Value(),
		HashedPassword: user.GetHashedPassword().Value(),
	}
	if err := ur.querier.InsertUser(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	u, err := ur.querier.FindUserByEmail(ctx, email.Value())
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

func (ur *userRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	u, err := ur.querier.FindUserById(ctx, id)
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
	us, err := ur.querier.FetchAllUser(ctx)
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
	params := UpdateUserParams{
		Name:  user.GetName(),
		Email: user.GetEmail().Value(),
		ID:    user.GetID(),
	}
	if err := ur.querier.UpdateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Delete(ctx context.Context, user *user.User) error {
	if err := ur.querier.DeleteUser(ctx, user.GetID()); err != nil {
		return err
	}
	return nil
}
