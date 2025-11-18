package models

// MatchedChar (เหมือนเดิม)
type MatchedChar struct {
	Character string `json:"character"`
	Value     int    `json:"value"`
}

// 🚀 [เปลี่ยนชื่อ] เปลี่ยน CalculationResult เป็น AstrologySet
// นี่คือชุดผลลัพธ์สำหรับ 1 ประเภท (เช่น เลขศาสตร์)
type AstrologySet struct {
	MatchedChars []MatchedChar `json:"matched_chars"`
	TotalSum     int           `json:"total_sum"`
}

// 🚀 [ใหม่] นี่คือ Model ผลลัพธ์ใหม่ที่จะส่งกลับไป
// มันจะห่อหุ้มผลลัพธ์ 2 ชุดไว้ด้วยกัน
type AnalysisResult struct {
	SatNum AstrologySet `json:"sat_num"`
	ShaNum AstrologySet `json:"sha_num"`
}

// AstrologyRequest (เหมือนเดิม)
type AstrologyRequest struct {
	Name string `json:"name"`
}
