package service

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.com/kovarniykrab/servermain/internel/domain"
)

func (app *App) GetUser(ctx context.Context, userID int) (domain.User, error) {
	u, e := app.Database.GetUserByID(ctx, userID)
	if e != nil {
		return domain.User{}, e
	}

	return u, nil
}

func (app *App) SearchUser(ctx context.Context, user domain.UserSearchForm) (domain.Users, error) {
	u, _, e := app.Database.UsersSearch(ctx, user)
	if e != nil {
		return domain.Users{}, e
	}

	return u, nil
}

func (app *App) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	u, err := app.Database.UpdateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (app *App) UpdateUser(ctx context.Context, id int, userForm domain.UserForm) (domain.User, error) {

	old, err := app.Database.GetUserByID(ctx, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return domain.User{}, domain.NotFoundErr(err)
		}
		return domain.User{}, err
	}

	old.UserForm = userForm
	u, err := app.Database.UpdateUser(ctx, old)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (app *App) DeleteUser(ctx context.Context, id int, userForm domain.UserForm) (domain.User, error) {

	_, err := app.Database.GetUserByID(ctx, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user with id %d not found", id)
		}
		return domain.User{}, fmt.Errorf("failed to check user existence: %w", err)
	}

	err = app.Database.DeleteUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to delete user: %v", err)
	}
	return domain.User{}, err
}
