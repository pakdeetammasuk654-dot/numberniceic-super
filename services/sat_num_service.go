package services

import (
	"log"
	"numberniceic/models"
	"numberniceic/repository"
	"strconv"
)

// SatNumService (Interface) (เหมือนเดิม)
type SatNumService interface {
	GetAllSatNums() ([]models.SatNum, error)
	CalculateNameAstrology(name string) (models.AnalysisResult, error)
}

// Struct (เหมือนเดิม)
type satNumService struct {
	satRepo repository.SatNumRepository
	shaRepo repository.ShaNumRepository
}

// New Service (เหมือนเดิม)
func NewSatNumService(satRepo repository.SatNumRepository, shaRepo repository.ShaNumRepository) SatNumService {
	return &satNumService{
		satRepo: satRepo,
		shaRepo: shaRepo,
	}
}

// 🚀 --- [เราจะแก้ไขฟังก์ชันนี้] --- 🚀
func (s *satNumService) CalculateNameAstrology(name string) (models.AnalysisResult, error) {

	// ---------------------------------------------------------------
	// 🚀 [เพิ่มส่วนนี้เพื่อ TEST] เราจะบังคับให้ Return ค่า Test
	//    เพื่อดูว่า "ั" ที่ถูกต้อง จะแสดงผลเพี้ยนหรือไม่
	// ---------------------------------------------------------------
	if name == "TestEncoding" { // 👈 [1] ลองพิมพ์ "TestEncoding" ในหน้าเว็บ
		log.Println("--- DEBUG: Bypassing database for Encoding Test ---")

		testResult := models.AnalysisResult{
			SatNum: models.AstrologySet{
				MatchedChars: []models.MatchedChar{
					{Character: "อ", Value: 1},
					{Character: "ั", Value: 4}, // 👈 [2] เรา Hardcode "ั" ที่นี่
					{Character: "ญ", Value: 5},
				},
				TotalSum: 10,
			},
			ShaNum: models.AstrologySet{ // 👈 (พลังเงา ก็ส่งค่าเปล่าๆไปก่อน)
				MatchedChars: []models.MatchedChar{},
				TotalSum:     0,
			},
		}
		return testResult, nil // 👈 [3] ส่งค่า Test กลับไปเลย
	}
	// --- จบส่วน TEST ---
	// (ถ้าพิมพ์ชื่ออื่นที่ไม่ใช่ "TestEncoding" มันจะทำงานตามปกติ)

	// --- 1. ดึงข้อมูล (SatNum) --- (โค้ดเดิม)
	allSatNums, err := s.satRepo.GetAllSatNums()
	if err != nil {
		return models.AnalysisResult{}, err
	}
	satMap := make(map[string]int)
	for _, satNum := range allSatNums {
		if satNum.CharKey != nil && satNum.SatValue != nil {
			val, err := strconv.Atoi(*satNum.SatValue)
			if err != nil {
				log.Printf("Warning: Skipping invalid SatValue for key %s: %v", *satNum.CharKey, err)
				continue
			}
			satMap[*satNum.CharKey] = val
		}
	}

	// --- 2. ดึงข้อมูล (ShaNum) --- (โค้ดเดิม)
	allShaNums, err := s.shaRepo.GetAllShaNums()
	if err != nil {
		return models.AnalysisResult{}, err
	}
	shaMap := make(map[string]int)
	for _, shaNum := range allShaNums {
		if shaNum.CharKey != nil && shaNum.ShaValue != nil {
			shaMap[*shaNum.CharKey] = int(*shaNum.ShaValue)
		}
	}

	// --- 3. เตรียมตัวแปร --- (โค้ดเดิม)
	var satTotalSum int = 0
	var satMatchedChars []models.MatchedChar
	var shaTotalSum int = 0
	var shaMatchedChars []models.MatchedChar

	// --- 4. วนลูป "ชื่อ" --- (โค้dเดิม)
	for _, charRune := range name {
		charStr := string(charRune)

		if val, ok := satMap[charStr]; ok {
			satTotalSum += val
			satMatchedChars = append(satMatchedChars, models.MatchedChar{
				Character: charStr,
				Value:     val,
			})
		}
		if val, ok := shaMap[charStr]; ok {
			shaTotalSum += val
			shaMatchedChars = append(shaMatchedChars, models.MatchedChar{
				Character: charStr,
				Value:     val,
			})
		}
	}

	// --- 5. สร้างผลลัพธ์สุดท้าย --- (โค้ดเดิม)
	result := models.AnalysisResult{
		SatNum: models.AstrologySet{
			MatchedChars: satMatchedChars,
			TotalSum:     satTotalSum,
		},
		ShaNum: models.AstrologySet{
			MatchedChars: shaMatchedChars,
			TotalSum:     shaTotalSum,
		},
	}

	return result, nil
}

// GetAllSatNums (เหมือนเดิม)
func (s *satNumService) GetAllSatNums() ([]models.SatNum, error) {
	// ... (โค้ดเดิม) ...
	satNums, err := s.satRepo.GetAllSatNums()
	if err != nil {
		return nil, err
	}
	return satNums, nil
}
