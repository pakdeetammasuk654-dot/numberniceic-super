package services

import (
	"fmt"
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

// Struct (เหมือนเดิม)
type analysisService struct {
	satRepo   repository.SatNumRepository
	shaRepo   repository.ShaNumRepository
	kakisRepo repository.KakisDayRepository
	numRepo   repository.NumberRepository
}

// New Service (เหมือนเดิม)
func NewAnalysisService(
	satRepo repository.SatNumRepository,
	shaRepo repository.ShaNumRepository,
	kakisRepo repository.KakisDayRepository,
	numRepo repository.NumberRepository,
) AnalysisService {
	return &analysisService{
		satRepo:   satRepo,
		shaRepo:   shaRepo,
		kakisRepo: kakisRepo,
		numRepo:   numRepo,
	}
}

// getMeaningsForSum (Helper) (เหมือนเดิม)
func (s *analysisService) getMeaningsForSum(sum int) []models.Number {
	var pairStrings []string

	if sum < 10 {
		pairStrings = append(pairStrings, fmt.Sprintf("0%d", sum))
	} else if sum > 99 {
		sumStr := fmt.Sprintf("%d", sum)
		pairStrings = append(pairStrings, sumStr[0:2])
		pairStrings = append(pairStrings, sumStr[len(sumStr)-2:])
	} else {
		pairStrings = append(pairStrings, fmt.Sprintf("%d", sum))
	}

	var meanings []models.Number
	for _, pair := range pairStrings {
		numberMeaning, err := s.numRepo.GetByPairNumber(pair)
		if err != nil {
			log.Printf("Warning: No meaning found for pairnumber %s: %v", pair, err)
			continue
		}
		meanings = append(meanings, numberMeaning)
	}
	return meanings
}

// 🚀 [ใหม่] ฟังก์ชัน Helper สำหรับคำนวณคะแนนรวม
// (ย้าย Logic มาจาก analyze_name.gohtml)
func (s *analysisService) calculateCombinedScores(satMeanings []models.Number, shaMeanings []models.Number) models.ScoreSummary {
	var goodScore int = 0
	var badScore int = 0 // เก็บเป็นค่าลบสะสม

	// 1. รวม Array 2 ชุด
	allMeanings := append(satMeanings, shaMeanings...)

	// 2. วนลูปเพื่อรวมคะแนน (ใช้ pairpoint)
	for _, m := range allMeanings {
		var p int = 0
		if m.PairPoint != nil {
			p = int(*m.PairPoint)
		}

		if p > 0 {
			goodScore += p
		} else if p < 0 {
			badScore += p // บวกค่าลบ
		}
	}

	// 3. ทำให้คะแนนร้ายเป็นบวก (Math.abs)
	var absBadScore int = badScore
	if absBadScore < 0 {
		absBadScore = -absBadScore
	}

	// 4. คืนค่า Struct ใหม่
	return models.ScoreSummary{
		GoodScore:  goodScore,
		BadScore:   absBadScore,
		TotalScore: goodScore + badScore, // (คะแนนดี + คะแนนร้าย(ที่เป็นค่าลบ)) = ผลรวมสุทธิ
	}
}

// 🚀 [อัปเดต] Logic การคำนวณ (CalculateNameAstrology)
func (s *analysisService) CalculateNameAstrology(name string, day string) (models.AnalysisResult, error) {

	// --- 1. ดึงข้อมูล (SatNum) --- (เหมือนเดิม)
	allSatNums, err := s.satRepo.GetAllSatNums()
	if err != nil {
		return models.AnalysisResult{}, err
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
	if err != nil {
		return models.AnalysisResult{}, err
	}
	shaMap := make(map[string]int)
	for _, shaNum := range allShaNums {
		if shaNum.CharKey != nil && shaNum.ShaValue != nil {
			shaMap[*shaNum.CharKey] = int(*shaNum.ShaValue)
		}
	}

	// --- 3. ดึงข้อมูล (Kakis) --- (เหมือนเดิม)
	kakisChars, err := s.kakisRepo.GetKakisByDay(day)
	if err != nil {
		return models.AnalysisResult{}, err
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
		if val, ok := satMap[charStr]; ok {
			satTotalSum += val
			satMatchedChars = append(satMatchedChars, models.MatchedChar{Character: charStr, Value: val})
		}
		if val, ok := shaMap[charStr]; ok {
			shaTotalSum += val
			shaMatchedChars = append(shaMatchedChars, models.MatchedChar{Character: charStr, Value: val})
		}
		if _, ok := kakisMap[charStr]; ok {
			// (ตรวจสอบว่ายังไม่มีอักษรนี้)
			found := false
			for _, k := range kakisFound {
				if k == charStr {
					found = true
					break
				}
			}
			if !found {
				kakisFound = append(kakisFound, charStr)
			}
		}
	}

	// --- 6. 🚀 [อัปเดต] สร้างผลลัพธ์สุดท้าย ---

	// (เหมือนเดิม) เรียก Helper เพื่อหาความหมาย
	satMeanings := s.getMeaningsForSum(satTotalSum)
	shaMeanings := s.getMeaningsForSum(shaTotalSum)

	// 🚀 [ใหม่] เรียกใช้ Helper เพื่อคำนวณคะแนน
	combinedScores := s.calculateCombinedScores(satMeanings, shaMeanings)

	result := models.AnalysisResult{
		SatNum: models.AstrologySet{
			MatchedChars: satMatchedChars,
			TotalSum:     satTotalSum,
			SumMeanings:  satMeanings,
		},
		ShaNum: models.AstrologySet{
			MatchedChars: shaMatchedChars,
			TotalSum:     shaTotalSum,
			SumMeanings:  shaMeanings,
		},
		KakisFound: kakisFound,

		// 🚀 [ใหม่] เพิ่ม field นี้เข้าไปในผลลัพธ์
		CombinedScoreSummary: combinedScores,
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
