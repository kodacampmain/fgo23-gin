package repositories

import (
	"context"
	"fgo23-gin/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PingRepository struct {
	db *pgxpool.Pool
}

func NewPingRepository(db *pgxpool.Pool) *PingRepository {
	return &PingRepository{db: db}
}

func (p *PingRepository) GetStudents(c context.Context) ([]models.Student, error) {
	query := "SELECT id, name FROM students"
	rows, err := p.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.Id, &student.Name); err != nil {
			return nil, err
		}
		result = append(result, student)
	}
	return result, nil
}
