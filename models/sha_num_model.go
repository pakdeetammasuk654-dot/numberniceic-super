package models

type ShaNum struct {
	CharKey  *string `json:"char_key,omitempty"`
	ShaValue *int32  `json:"sha_value,omitempty"` // 👈 ระวัง: Table DDL ของคุณคือ integer
}
