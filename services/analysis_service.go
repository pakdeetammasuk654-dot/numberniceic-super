package services

import (
	"log"
	"numberniceic/models"
	"numberniceic/repository"
	"strconv"
)

// 🚀 [อัปเดต] Interface
type AnalysisService interface {
	GetAllSatNums() ([]models.SatNum, error)
	// 👈 เพิ่ม day string
	CalculateNameAstrology(name string, day string) (models.AnalysisResult, error)
}

// 🚀 [อัปเดต] Struct
type analysisService struct {
	satRepo   repository.SatNumRepository
	shaRepo   repository.ShaNumRepository
	kakisRepo repository.KakisDayRepository // 👈 เพิ่ม repo ใหม่
}

// 🚀 [อัปเดต] New Service
func NewAnalysisService(
	satRepo repository.SatNumRepository,
	shaRepo repository.ShaNumRepository,
	kakisRepo repository.KakisDayRepository, // 👈 เพิ่ม repo ใหม่
) AnalysisService {
	return &analysisService{
		satRepo:   satRepo,
		shaRepo:   shaRepo,
		kakisRepo: kakisRepo, // 👈 เพิ่ม repo ใหม่
	}
}

// 🚀 [อัปเดต] Logic การคำนวณ
func (s *analysisService) CalculateNameAstrology(name string, day string) (models.AnalysisResult, error) {

	// --- 1. ดึงข้อมูล (SatNum) --- (เหมือนเดิม)
	allSatNums, err := s.satRepo.GetAllSatNums()
	if err != nil { /* ... */
	}
	satMap := make(map[string]int)
	for _, satNum := range allSatNums {
		if satNum.CharKey != nil && satNum.SatValue != nil {
			val, _ := strconv.Atoi(*satNum.SatValue)
			satMap[*satNum.CharKey] = val
		}
	}

	// --- 2. ดึงข้อมูล (ShaNum) --- (เหมือนเดิม)
	allShaNums, err := s.shaRepo.GetAllShaNums()
	if err != nil { /* ... */
	}
	shaMap := make(map[string]int)
	for _, shaNum := range allShaNums {
		if shaNum.CharKey != nil && shaNum.ShaValue != nil {
			shaMap[*shaNum.CharKey] = int(*shaNum.ShaValue)
		}
	}

	// --- 3. 🚀 [ใหม่] ดึงข้อมูล (Kakis) ---
	kakisChars, err := s.kakisRepo.GetKakisByDay(day)
	if err != nil {
		// (ถ้าหาไม่เจอ ก็แค่ Log ไว้ แต่ไม่ควรหยุดการทำงาน)
		log.Printf("Warning: Could not retrieve Kakis for day %s: %v", day, err)
	}
	// สร้าง Map เพื่อค้นหา Kakis ได้เร็ว
	kakisMap := make(map[string]bool)
	for _, char := range kakisChars {
		kakisMap[char] = true
	}

	// --- 4. เตรียมตัวแปร (อัปเดต) ---
	var satTotalSum int = 0
	var satMatchedChars []models.MatchedChar
	var shaTotalSum int = 0
	var shaMatchedChars []models.MatchedChar
	var kakisFound []string // 👈 [ใหม่]

	// --- 5. วนลูป "ชื่อ" (อัปเดต) ---
	for _, charRune := range name {
		charStr := string(charRune)

		// 5a. Check SatNum (เหมือนเดิม)
		if val, ok := satMap[charStr]; ok {
			satTotalSum += val
			satMatchedChars = append(satMatchedChars, models.MatchedChar{
				Character: charStr,
				Value:     val,
			})
		}

		// 5b. Check ShaNum (เหมือนเดิม)
		if val, ok := shaMap[charStr]; ok {
			shaTotalSum += val
			shaMatchedChars = append(shaMatchedChars, models.MatchedChar{
				Character: charStr,
				Value:     val,
			})
		}

		// 5c. 🚀 [ใหม่] Check Kakis
		if _, ok := kakisMap[charStr]; ok {
			kakisFound = append(kakisFound, charStr)
		}
	}

	// --- 6. สร้างผลลัพธ์สุดท้าย (อัปเดต) ---
	result := models.AnalysisResult{
		SatNum: models.AstrologySet{
			MatchedChars: satMatchedChars,
			TotalSum:     satTotalSum,
		},
		ShaNum: models.AstrologySet{
			MatchedChars: shaMatchedChars,
			TotalSum:     shaTotalSum,
		},
		KakisFound: kakisFound, // 👈 [ใหม่]
	}

	return result, nil
}

// GetAllSatNums (เหมือนเดิม)
func (s *analysisService) GetAllSatNums() ([]models.SatNum, error) {
	satNums, err := s.satRepo.GetAllSatNums()
	if err != nil {
		return nil, err
	}
	return satNums, nil
}
