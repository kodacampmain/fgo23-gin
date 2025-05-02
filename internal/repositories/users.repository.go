package repositories

import (
	"context"
	"fgo23-gin/internal/models"
	"fmt"

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

func (u *UserRepository) UpdateStudents(c context.Context, id int, body models.Student, imageUrl string) (models.Student, error) {
	// Creating Query (Query Builder)
	query := "UPDATE students SET "
	// UPDATE table SET field1=value1, field2=value2 where field=value
	values := []any{}
	if body.Name != "" {
		query += fmt.Sprintf("name=$%d", len(values)+1)
		values = append(values, body.Name)
	}
	if body.Image != "" {
		if len(values) > 0 {
			query += ", "
		}
		query += fmt.Sprintf("images=$%d", len(values)+1)
		values = append(values, imageUrl)
	}
	query += fmt.Sprintf(" WHERE id=$%d", len(values)+1)
	values = append(values, id)
	query += " RETURNING id,name,images"
	// END Creating Query
	var result models.Student
	if err := u.db.QueryRow(c, query, values...).Scan(&result.Id, &result.Name, &result.Image); err != nil {
		return models.Student{}, err
	}
	return result, nil
}
