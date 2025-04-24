package repositories

import (
	"context"
	"fgo23-gin/internal/models"
	"fgo23-gin/pkg"
)

type PingRepository struct{}

var PingRepo *PingRepository

func NewPingRepository() {
	PingRepo = &PingRepository{}
}

func (p *PingRepository) GetStudents(c context.Context) ([]models.Student, error) {
	query := "SELECT id, name FROM students"
	rows, err := pkg.DB.Query(c, query)
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
