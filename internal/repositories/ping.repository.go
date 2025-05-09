package repositories

import (
	"context"
	"encoding/json"
	"fgo23-gin/internal/models"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PingRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewPingRepository(db *pgxpool.Pool, rdb *redis.Client) *PingRepository {
	return &PingRepository{db: db, rdb: rdb}
}

func (p *PingRepository) GetStudents(c context.Context) ([]models.Student, error) {
	// cek redis terlebih dahulu, jika ada nilainya, maka gunakan nilai dari redis
	redisKey := "students"
	cache, err := p.rdb.Get(c, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("\nkey %s does not exist\n", redisKey)
		} else {
			log.Println("Redis not working")
		}
	} else {
		var students []models.Student
		if err := json.Unmarshal([]byte(cache), &students); err != nil {
			return nil, err
		}
		if len(students) > 0 {
			return students, nil
		}
	}
	// jika di redis tidak ada, maka ambil data dari db
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
	// data baru yang diambil masukkan ke dalam redis
	res, err := json.Marshal(result)
	if err != nil {
		log.Println("[DEBUG] marshal", err.Error())
	}
	if err := p.rdb.Set(c, redisKey, string(res), time.Minute*5).Err(); err != nil {
		log.Println("[DEBUG] redis set", err.Error())
	}
	return result, nil
}
