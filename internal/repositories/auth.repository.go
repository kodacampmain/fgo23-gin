package repositories

import (
	"context"
	"fgo23-gin/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) AddNewUser(context context.Context, name string, hashedPass string) (pgconn.CommandTag, error) {
	query := "INSERT INTO students (name, password) VALUES ($1, $2)"
	values := []any{name, hashedPass}

	// cmd, err := a.db.Exec(context, query, values...)
	// if err != nil {
	// 	return pgconn.CommandTag{}, err
	// }
	// return cmd, nil
	return a.db.Exec(context, query, values...)
}

func (a *AuthRepo) GetUserData(context context.Context, name string) (models.Student, error) {
	query := "select id, name, password from students where name = $1"
	values := []any{name}

	var result models.Student
	if err := a.db.QueryRow(context, query, values...).Scan(&result.Id, &result.Name, &result.Password); err != nil {
		return models.Student{}, err
	}
	return result, nil
}
