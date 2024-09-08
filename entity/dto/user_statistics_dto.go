package dto

type UserStatisticsDto struct {
	Type  int     `json:"type"`
	Value float32 `json:"value"`
	Num   int     `json:"num"`
}
