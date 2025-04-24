package repositories

import (
	"context"
	"fgo23-gin/internal/models"
	"fgo23-gin/pkg"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct{}

var UserRepo *UserRepository

func NewUserRepository() {
	UserRepo = &UserRepository{}
}

func (u *UserRepository) FindEmployeeById(c context.Context, id int, name string) (models.Employee, error) {
	query := "SELECT id,name,salary FROM employee WHERE id=$1 AND name=$2"
	values := []any{id, name}
	var result models.Employee
	if err := pkg.DB.QueryRow(c, query, values...).Scan(&result.Id, &result.Name, &result.Salary); err != nil {
		return models.Employee{}, err
	}
	return result, nil
}

func (u *UserRepository) CreateNewEmployee(c context.Context, newEmployee models.Employee) (pgconn.CommandTag, error) {
	query := "INSERT INTO employee (name, salary, city) VALUES ($1, $2, $3)"
	values := []any{newEmployee.Name, newEmployee.Salary, newEmployee.City}
	cmd, err := pkg.DB.Exec(c, query, values...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return cmd, nil
}
