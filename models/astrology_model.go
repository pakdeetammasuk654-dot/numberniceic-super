package models

// MatchedChar (เหมือนเดิม)
type MatchedChar struct {
	Character string `json:"character"`
	Value     int    `json:"value"`
}

// AstrologySet (อัปเดต)
type AstrologySet struct {
	MatchedChars []MatchedChar `json:"matched_chars"`
	TotalSum     int           `json:"total_sum"`

	// 🚀 [ใหม่] เราจะเก็บ "ความหมาย" ของผลรวม (จากตาราง numbers)
	// (ใช้ []Number เพราะผลรวม 100 อาจได้ 2 ความหมาย)
	SumMeanings []Number `json:"sum_meanings,omitempty"`
}

// AnalysisResult (เหมือนเดิม)
type AnalysisResult struct {
	SatNum     AstrologySet `json:"sat_num"`
	ShaNum     AstrologySet `json:"sha_num"`
	KakisFound []string     `json:"kakis_found,omitempty"`
}

// AstrologyRequest (เหมือนเดิม)
type AstrologyRequest struct {
	Name string `json:"name"`
	Day  string `json:"day"`
}
