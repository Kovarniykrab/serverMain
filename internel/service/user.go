package service

import (
	"context"
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
