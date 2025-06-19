package database

import (
	"context"
	"github.com/lib/pq"
	"gitlab.com/kovarniykrab/servermain/internel/domain"
)

func (db *Service) UsersSearch(ctx context.Context, form domain.UserSearchForm) (domain.Users, int, error) {
	dm := domain.Users{}

	q := db.db.NewSelect().Model(&dm)

	if len(form.IDs) > 0 {
		q.Where("id = ANY(?)",
			pq.Array(form.IDs))
	}

	if len(form.ExcludedIDs) > 0 {
		q.Where("NOT (id = ANY(?))",
			pq.Array(form.ExcludedIDs))
	}

	if len(form.IDs) > 0 {
		q.Where("id = ANY(?)",
			pq.Array(form.IDs))
	}

	if form.UserName != "" {
		q.Where("user_name = ?", form.UserName)
	}

	cnt, err := q.Offset((form.Page - 1) * form.Limit).
		Limit(form.Limit).ScanAndCount(ctx)

	if err != nil {
		return dm, cnt, domain.IntErr(err)
	}

	return dm, cnt, nil
}

func (db *Service) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	md := domain.User{}

	if err := db.db.NewSelect().Model(&md).Where("id = ?", id).Scan(ctx); err != nil {
		return md, err
	}

	return md, nil
}

func (db *Service) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if _, err := db.db.NewInsert().
		Model(&user).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx); err != nil {
		return user, err
	}

	return user, nil
}

func (db *Service) DeleteUser(ctx context.Context, id int) error {

	if _, err := db.db.NewDelete().Model(&domain.User{ID: id}).
		WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}
