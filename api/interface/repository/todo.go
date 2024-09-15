package repository

import (
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"

	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

type TodoRepository struct {
	Conn *sqlx.DB
}

func NewTodoRepository(conn *sqlx.DB) repository.TodoRepositoryInterface {
	return &TodoRepository{
		Conn: conn,
	}
}

func (r *TodoRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*entity.Todo, err error) {
    err = r.Conn.SelectContext(ctx, &result, query, args...)
    if err != nil {
        slog.Error("Todo fetch query", slog.String("error", err.Error()))
        return nil, err
    }
    return result, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *entity.Todo) (err error) {
	query := `INSERT INTO todo (title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	err = withTransaction(ctx, r.Conn, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, query, todo)
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

func (r *TodoRepository) GetList(ctx context.Context) (todos []*entity.Todo, err error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todo`

	res, err := r.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TodoRepository) GetByID(ctx context.Context, id string) (todo *entity.Todo, err error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todo WHERE id = ?`

	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		todo = list[0]
	} else {
		return nil, entity.ErrNotFound
	}

	return todo, nil
}

func (r *TodoRepository) Update(ctx context.Context, todo *entity.Todo) (err error) {
	query := `UPDATE todo SET title = :title, description = :description, completed = :completed, updated_at = updated_at WHERE id = :id`

	err = withTransaction(ctx, r.Conn, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, query, todo)
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

func (r *TodoRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM todo WHERE id = ?`

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