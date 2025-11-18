package repository

import (
	"database/sql"
	"errors"
	"log"
	"numberniceic/models"
)

// [ใหม่] สร้าง error มาตรฐานสำหรับ "ไม่พบข้อมูล"
// นี่เป็น Pattern ที่ดีมาก เพื่อให้ Service Layer ไม่ต้อง import "database/sql"
var ErrNotFound = errors.New("number not found")

type NumberRepository interface {
	GetAllNumbers() ([]models.Number, error)
	GetByPairNumber(pairNumber string) (models.Number, error)
}

type numberRepository struct {
	DB *sql.DB
}

func NewNumberRepository(db *sql.DB) NumberRepository {
	return &numberRepository{DB: db}
}

func (r *numberRepository) GetByPairNumber(pairNumber string) (models.Number, error) {
	// ใช้ QueryRow เพราะเราคาดหวังผลลัพธ์เดียว
	query := `
		SELECT 
			detail_vip, pairtype, pairnumber, 
			miracledetail, miracledesc, pairnumberid, pairpoint 
		FROM numbers
		WHERE pairnumber = $1
	`

	var n models.Number

	// Scan ข้อมูลเข้า Pointers
	err := r.DB.QueryRow(query, pairNumber).Scan(
		&n.DetailVip,
		&n.PairType,
		&n.PairNumber,
		&n.MiracleDetail,
		&n.MiracleDesc,
		&n.PairNumberID,
		&n.PairPoint,
	)

	if err != nil {
		// 3. [สำคัญ] แปลง error
		if errors.Is(err, sql.ErrNoRows) {
			// ถ้าไม่เจอ ให้คืนค่า error ที่เราสร้างไว้
			return models.Number{}, ErrNotFound
		}
		// ถ้าเป็น error อื่น (เช่น DB พัง)
		log.Printf("Error scanning row: %v", err)
		return models.Number{}, err
	}

	return n, nil
}

func (r *numberRepository) GetAllNumbers() ([]models.Number, error) {
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
