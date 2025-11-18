package repository

import (
	"database/sql"
	"log"
	"numberniceic/models"
)

type SatNumRepository interface {
	GetAllSatNums() ([]models.SatNum, error)
}

type satNumRepository struct {
	DB *sql.DB
}

func (s *satNumRepository) GetAllSatNums() ([]models.SatNum, error) {
	rows, err := s.DB.Query(`SELECT char_key, sat_value FROM sat_nums`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.SatNum

	for rows.Next() {
		var n models.SatNum
		err := rows.Scan(&n.CharKey, &n.SatValue)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		results = append(results, n)
	}
	return results, nil

}

func NewSatNumRepository(db *sql.DB) SatNumRepository {
	return &satNumRepository{DB: db}
}
