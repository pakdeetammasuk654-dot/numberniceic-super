package services

import (
	"fmt" // 👈 [ใหม่]
	"log"
	"numberniceic/models"
	"numberniceic/repository"
	"strconv"
)

// AnalysisService (Interface) (เหมือนเดิม)
type AnalysisService interface {
	GetAllSatNums() ([]models.SatNum, error)
	CalculateNameAstrology(name string, day string) (models.AnalysisResult, error)
}

// 🚀 [อัปเดต] Struct
type analysisService struct {
	satRepo   repository.SatNumRepository
	shaRepo   repository.ShaNumRepository
	kakisRepo repository.KakisDayRepository
	numRepo   repository.NumberRepository // 👈 [ใหม่]
}

// 🚀 [อัปเดต] New Service
func NewAnalysisService(
	satRepo repository.SatNumRepository,
	shaRepo repository.ShaNumRepository,
	kakisRepo repository.KakisDayRepository,
	numRepo repository.NumberRepository, // 👈 [ใหม่]
) AnalysisService {
	return &analysisService{
		satRepo:   satRepo,
		shaRepo:   shaRepo,
		kakisRepo: kakisRepo,
		numRepo:   numRepo, // 👈 [ใหม่]
	}
}

// 🚀 [ใหม่] ฟังก์ชัน Helper สำหรับแปลงผลรวมเป็นความหมาย
func (s *analysisService) getMeaningsForSum(sum int) []models.Number {
	var pairStrings []string

	if sum < 10 {
		// 1. กฎหลักหน่วย (เช่น 8 -> "08")
		pairStrings = append(pairStrings, fmt.Sprintf("0%d", sum))
	} else if sum > 99 {
		// 2. กฎหลักร้อย (เช่น 100 -> "10", "00" | 123 -> "12", "23")
		sumStr := fmt.Sprintf("%d", sum)                          // เช่น "123"
		pairStrings = append(pairStrings, sumStr[0:2])            // "12"
		pairStrings = append(pairStrings, sumStr[len(sumStr)-2:]) // "23"
	} else {
		// 3. กฎหลักสิบ (เช่น 45 -> "45")
		pairStrings = append(pairStrings, fmt.Sprintf("%d", sum))
	}

	// 4. ดึงความหมายจาก DB
	var meanings []models.Number
	for _, pair := range pairStrings {
		numberMeaning, err := s.numRepo.GetByPairNumber(pair)
		if err != nil {
			// ถ้าหาไม่เจอ ก็แค่ Log ไว้
			log.Printf("Warning: No meaning found for pairnumber %s: %v", pair, err)
			continue
		}
		meanings = append(meanings, numberMeaning)
	}
	return meanings
}

// 🚀 [อัปเดต] Logic การคำนวณ (CalculateNameAstrology)
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

	// --- 3. ดึงข้อมูล (Kakis) --- (เหมือนเดิม)
	kakisChars, err := s.kakisRepo.GetKakisByDay(day)
	if err != nil { /* ... */
	}
	kakisMap := make(map[string]bool)
	for _, char := range kakisChars {
		kakisMap[char] = true
	}

	// --- 4. เตรียมตัวแปร (เหมือนเดิม) ---
	var satTotalSum int = 0
	var satMatchedChars []models.MatchedChar
	var shaTotalSum int = 0
	var shaMatchedChars []models.MatchedChar
	var kakisFound []string

	// --- 5. วนลูป "ชื่อ" (เหมือนเดิม) ---
	for _, charRune := range name {
		charStr := string(charRune)
		if val, ok := satMap[charStr]; ok { /* ... */
			satTotalSum += val
			satMatchedChars = append(satMatchedChars, models.MatchedChar{Character: charStr, Value: val})
		}
		if val, ok := shaMap[charStr]; ok { /* ... */
			shaTotalSum += val
			shaMatchedChars = append(shaMatchedChars, models.MatchedChar{Character: charStr, Value: val})
		}
		if _, ok := kakisMap[charStr]; ok {
			kakisFound = append(kakisFound, charStr)
		}
	}

	// --- 6. 🚀 [อัปเดต] สร้างผลลัพธ์สุดท้าย ---

	// 👈 [ใหม่] เรียก Helper เพื่อหาความหมาย
	satMeanings := s.getMeaningsForSum(satTotalSum)
	shaMeanings := s.getMeaningsForSum(shaTotalSum)

	result := models.AnalysisResult{
		SatNum: models.AstrologySet{
			MatchedChars: satMatchedChars,
			TotalSum:     satTotalSum,
			SumMeanings:  satMeanings, // 👈 [ใหม่]
		},
		ShaNum: models.AstrologySet{
			MatchedChars: shaMatchedChars,
			TotalSum:     shaTotalSum,
			SumMeanings:  shaMeanings, // 👈 [ใหม่]
		},
		KakisFound: kakisFound,
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
