package models

// MatchedChar เก็บรายละเอียดของอักษรที่ตรงกับเลขศาสตร์
type MatchedChar struct {
	Character string `json:"character"`
	Value     int    `json:"value"`
}

// CalculationResult คือผลลัพธ์ทั้งหมดที่จะส่งกลับไป
type CalculationResult struct {
	MatchedChars []MatchedChar `json:"matched_chars"`
	TotalSum     int           `json:"total_sum"`
}

// AstrologyRequest คือโครงสร้าง JSON ที่เราคาดหวังจาก user
type AstrologyRequest struct {
	Name string `json:"name"`
}
