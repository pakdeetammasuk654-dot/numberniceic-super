package repository

import (
	"database/sql"
	"log"
	"numberniceic/models"
)

type NumberRepository struct {
	DB *sql.DB
}

func NewNumberRepository(db *sql.DB) *NumberRepository {
	return &NumberRepository{DB: db}
}

func (r *NumberRepository) GetAllNumbers() ([]models.Number, error) {
	rows, err := r.DB.Query(`
		SELECT 
			detail_vip, pairtype, pairnumber, 
			miracledetail, miracledesc, pairnumberid, pairpoint 
		FROM numbers
	`)
	if err != nil {
		log.Printf("Error querying numbers: %v", err)
		return nil, err // ส่ง error กลับไปให้ Handler จัดการ
	}
	defer rows.Close()

	var results []models.Number

	for rows.Next() {
		var n models.Number

		err := rows.Scan(
			&n.DetailVip,
			&n.PairType,
			&n.PairNumber,
			&n.MiracleDetail,
			&n.MiracleDesc,
			&n.PairNumberID,
			&n.PairPoint,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		results = append(results, n)
	}

	// คืนค่าข้อมูลที่ได้ และ nil (เพราะไม่มี error)
	return results, nil
}
