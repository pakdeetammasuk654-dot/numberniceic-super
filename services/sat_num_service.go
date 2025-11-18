package services

import (
	"log"
	"numberniceic/models"
	"numberniceic/repository"
	"strconv"
)

// SatNumService defines the interface for the sat_num service.
type SatNumService interface {
	GetAllSatNums() ([]models.SatNum, error)
	CalculateNameAstrology(name string) (models.CalculationResult, error)
}

type satNumService struct {
	repo repository.SatNumRepository
}

func (s *satNumService) CalculateNameAstrology(name string) (models.CalculationResult, error) {
	// 1. ดึงข้อมูลเลขศาสตร์ทั้งหมดจาก DB
	allSatNums, err := s.repo.GetAllSatNums()
	if err != nil {
		return models.CalculationResult{}, err
	}

	// 2. สร้าง "Map" เพื่อให้ค้นหาอักษรได้เร็ว
	// (แปลงจาก []models.SatNum เป็น map[string]int)
	satMap := make(map[string]int)
	for _, satNum := range allSatNums {
		if satNum.CharKey != nil && satNum.SatValue != nil {
			// 3. แปลงค่า SatValue (string) เป็น (int)
			val, err := strconv.Atoi(*satNum.SatValue)
			if err != nil {
				// ถ้าแปลงไม่ได้ ให้ข้ามไป (หรือ Log error ไว้)
				log.Printf("Warning: Skipping invalid SatValue for key %s: %v", *satNum.CharKey, err)
				continue
			}
			satMap[*satNum.CharKey] = val
		}
	}

	// 4. เตรียมตัวแปรสำหรับผลลัพธ์
	var totalSum int = 0
	var matchedChars []models.MatchedChar

	// 5. วนลูป "ชื่อ" ที่ส่งเข้ามา ทีละ "อักษร"
	// (การวนลูปแบบนี้รองรับภาษาไทยได้ถูกต้อง)
	for _, charRune := range name {
		charStr := string(charRune)

		// 6. ค้นหาอักษรใน Map ที่เราเตรียมไว้
		// val คือค่าตัวเลข, ok คือ bool (ว่าเจอหรือไม่)
		if val, ok := satMap[charStr]; ok {
			// ถ้าเจอ (ok == true)
			totalSum += val // เพิ่มค่าในผลรวม
			matchedChars = append(matchedChars, models.MatchedChar{
				Character: charStr,
				Value:     val,
			})
		}
		// ถ้าไม่เจอ ก็แค่ข้ามไป ไม่ต้องทำอะไร
	}

	// 7. สร้างผลลัพธ์สุดท้าย
	result := models.CalculationResult{
		MatchedChars: matchedChars,
		TotalSum:     totalSum,
	}

	return result, nil
}

// NewSatNumService creates a new SatNumService.
func NewSatNumService(repo repository.SatNumRepository) SatNumService {
	return &satNumService{repo: repo}
}

// GetAllSatNums retrieves all sat_num records.
func (s *satNumService) GetAllSatNums() ([]models.SatNum, error) {
	// ในอนาคต เราสามารถเพิ่ม business logic ตรงนี้ได้
	// แต่ตอนนี้ เราจะแค่ส่งต่อไปให้ repository
	satNums, err := s.repo.GetAllSatNums()
	if err != nil {
		return nil, err // ส่ง error กลับขึ้นไป
	}
	return satNums, nil
}
