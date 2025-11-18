package repository

import (
	"database/sql"
	"log"
)

type KakisDayRepository interface {
	// 👈 [สำคัญ] เราจะให้มันคืนค่า []string (รายการอักษร)
	GetKakisByDay(day string) ([]string, error)
}

type kakisDayRepository struct {
	DB *sql.DB
}

func NewKakisDayRepository(db *sql.DB) KakisDayRepository {
	return &kakisDayRepository{DB: db}
}

func (r *kakisDayRepository) GetKakisByDay(day string) ([]string, error) {
	// 🚀 [สำคัญ] เราใช้ TRIM() เพื่อลบช่องว่างที่อาจติดมากับข้อมูล (เช่น '    ก')
	query := `SELECT TRIM(kakis) FROM kakis_day WHERE day = $1`

	rows, err := r.DB.Query(query, day)
	if err != nil {
		log.Printf("Error querying kakis_day: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var kakisChar string
		if err := rows.Scan(&kakisChar); err != nil {
			log.Printf("Error scanning kakis_day row: %v", err)
			continue
		}
		results = append(results, kakisChar)
	}

	return results, nil
}
