package repository

import (
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"

	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

type UserRepository struct {
	Conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) repository.UserRepositoryInterface {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*entity.User, err error) {
	err = r.Conn.SelectContext(ctx, &result, query, args...)
	if err != nil {
		slog.Error("User fetch error", slog.String("error", err.Error()))
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (err error) {
	query := `INSERT INTO user (id, name, email, auth0_id, role, created_at, updated_at) VALUES (:id, :name, :email, :auth0_id, :role, :created_at, :updated_at)`

	err = withTransaction(ctx, r.Conn, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, query, user)
		if err != nil {
			return err
		}

		if err = checkRowsAffected(res); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *UserRepository) GetList(ctx context.Context) (users []*entity.User, err error) {
	query := `SELECT id, name, email, auth0_id, role, created_at, updated_at FROM user`
	res, err := r.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (user *entity.User, err error) {
	query := `SELECT id, name, email, auth0_id, role, created_at, updated_at FROM user WHERE id = ?`
	users, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return nil, entity.ErrNotFound
}

func (r *UserRepository) GetByAuth0ID(ctx context.Context, auth0ID string) (user *entity.User, err error) {
	query := `SELECT id, name, email, auth0_id, role, created_at, updated_at FROM user WHERE auth0_id = ?`
	users, err := r.fetch(ctx, query, auth0ID)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return nil, entity.ErrNotFound
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) (err error) {
	query := `UPDATE user SET name = :name, email = :email, auth0_id = :auth0_id, role = :role, created_at = :created_at, updated_at = :updated_at WHERE id = :id`

	err = withTransaction(ctx, r.Conn, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, query, user)
		if err != nil {
			return err
		}
		if err = checkRowsAffected(res); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM user WHERE id = ?`

	err = withTransaction(ctx, r.Conn, func(tx *sqlx.Tx) error {
		res, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
		if err = checkRowsAffected(res); err != nil {
			return err
		}
		return nil
	})
	return err
}
