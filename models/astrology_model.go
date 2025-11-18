package models

// MatchedChar (เหมือนเดิม)
type MatchedChar struct {
	Character string `json:"character"`
	Value     int    `json:"value"`
}

// AstrologySet (เหมือนเดิม)
type AstrologySet struct {
	MatchedChars []MatchedChar `json:"matched_chars"`
	TotalSum     int           `json:"total_sum"`
}

// AnalysisResult (อัปเดต)
type AnalysisResult struct {
	SatNum AstrologySet `json:"sat_num"`
	ShaNum AstrologySet `json:"sha_num"`
	// 🚀 [ใหม่] เพิ่มช่องสำหรับอักษรกาลกิณีที่พบ
	KakisFound []string `json:"kakis_found,omitempty"`
}

// AstrologyRequest (อัปเดต)
type AstrologyRequest struct {
	Name string `json:"name"`
	// 🚀 [ใหม่] เพิ่มช่องสำหรับรับ "วันเกิด"
	Day string `json:"day"`
}
