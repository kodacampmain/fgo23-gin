package repositories

import (
	"context"
	"fgo23-gin/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindEmployeeById(c context.Context, id int, name string) (models.Employee, error) {
	query := "SELECT id,name,salary FROM employee WHERE id=$1 AND name=$2"
	values := []any{id, name}
	var result models.Employee
	if err := u.db.QueryRow(c, query, values...).Scan(&result.Id, &result.Name, &result.Salary); err != nil {
		return models.Employee{}, err
	}
	return result, nil
}

func (u *UserRepository) CreateNewEmployee(c context.Context, newEmployee models.Employee) (pgconn.CommandTag, error) {
	query := "INSERT INTO employee (name, salary, city) VALUES ($1, $2, $3)"
	values := []any{newEmployee.Name, newEmployee.Salary, newEmployee.City}
	cmd, err := u.db.Exec(c, query, values...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return cmd, nil
}
