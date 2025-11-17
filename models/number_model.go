package models

type Number struct {
	DetailVip     *string `json:"detail_vip,omitempty"`
	PairType      *string `json:"pairtype,omitempty"`
	PairNumber    *string `json:"pairnumber,omitempty"`
	MiracleDetail *string `json:"miracledetail,omitempty"`
	MiracleDesc   *string `json:"miracledesc,omitempty"`
	PairNumberID  *int32  `json:"pairnumberid,omitempty"`
	PairPoint     *int32  `json:"pairpoint,omitempty"`
}
