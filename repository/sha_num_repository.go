package repository

import (
	"database/sql"
	"log"
	"numberniceic/models"
)

type ShaNumRepository interface {
	GetAllShaNums() ([]models.ShaNum, error)
}

type shaNumRepository struct {
	DB *sql.DB
}

func NewShaNumRepository(db *sql.DB) ShaNumRepository {
	return &shaNumRepository{DB: db}
}

func (s *shaNumRepository) GetAllShaNums() ([]models.ShaNum, error) {
	rows, err := s.DB.Query(`SELECT char_key, sha_value FROM sha_nums`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ShaNum

	for rows.Next() {
		var n models.ShaNum
		// 👈 ระวัง: ตรวจสอบว่า Scan ถูกต้องตาม Model
		err := rows.Scan(&n.CharKey, &n.ShaValue)
		if err != nil {
			log.Printf("Error scanning sha_num row: %v", err)
			continue
		}
		results = append(results, n)
	}
	return results, nil
}
