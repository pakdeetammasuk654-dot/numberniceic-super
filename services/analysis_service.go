package services

import (
	"log"
	"numberniceic/models"
	"numberniceic/repository"
	"strconv"
)

// 🚀 [เปลี่ยน] เปลี่ยน Interface
type AnalysisService interface {
	GetAllSatNums() ([]models.SatNum, error)
	CalculateNameAstrology(name string) (models.AnalysisResult, error)
}

// 🚀 [เปลี่ยน] เปลี่ยน struct
type analysisService struct {
	satRepo repository.SatNumRepository
	shaRepo repository.ShaNumRepository
}

// 🚀 [เปลี่ยน] เปลี่ยนชื่อฟังก์ชัน New
func NewAnalysisService(satRepo repository.SatNumRepository, shaRepo repository.ShaNumRepository) AnalysisService {
	return &analysisService{
		satRepo: satRepo,
		shaRepo: shaRepo,
	}
}

// (ฟังก์ชัน CalculateNameAstrology เหมือนเดิม)
// (เราจะลบโค้ด TestEncoding ออก)
func (s *analysisService) CalculateNameAstrology(name string) (models.AnalysisResult, error) {

	// --- 1. ดึงข้อมูล (SatNum) ---
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

	// --- 2. ดึงข้อมูล (ShaNum) ---
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

	// --- 3. เตรียมตัวแปร ---
	var satTotalSum int = 0
	var satMatchedChars []models.MatchedChar
	var shaTotalSum int = 0
	var shaMatchedChars []models.MatchedChar

	// --- 4. วนลูป "ชื่อ" ---
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

	// --- 5. สร้างผลลัพธ์สุดท้าย ---
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

// (ฟังก์ชัน GetAllSatNums เหมือนเดิม)
func (s *analysisService) GetAllSatNums() ([]models.SatNum, error) {
	satNums, err := s.satRepo.GetAllSatNums()
	if err != nil {
		return nil, err
	}
	return satNums, nil
}
