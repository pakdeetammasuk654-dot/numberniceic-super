package models

// 🚀 [ใหม่] สร้าง Struct สำหรับสรุปคะแนน
type ScoreSummary struct {
	GoodScore  int `json:"good_score"`
	BadScore   int `json:"bad_score"`   // จะเป็นค่าบวก (Math.abs)
	TotalScore int `json:"total_score"` // คะแนนสุทธิ (ดี - ร้าย)
}

// MatchedChar (เหมือนเดิม)
type MatchedChar struct {
	Character string `json:"character"`
	Value     int    `json:"value"`
}

// AstrologySet (เหมือนเดิม)
type AstrologySet struct {
	MatchedChars []MatchedChar `json:"matched_chars"`
	TotalSum     int           `json:"total_sum"`

	// 🚀 [แก้ไข] field นี้มีอยู่แล้ว
	SumMeanings []Number `json:"sum_meanings,omitempty"`
}

// AnalysisResult (🚀 [อัปเดต] เพิ่ม field ใหม่)
type AnalysisResult struct {
	SatNum     AstrologySet `json:"sat_num"`
	ShaNum     AstrologySet `json:"sha_num"`
	KakisFound []string     `json:"kakis_found,omitempty"`

	// 🚀 [ใหม่] เพิ่มผลสรุปคะแนนที่คำนวณแล้ว
	CombinedScoreSummary ScoreSummary `json:"combined_score_summary"`
}

// AstrologyRequest (เหมือนเดิม)
type AstrologyRequest struct {
	Name string `json:"name"`
	Day  string `json:"day"`
}
