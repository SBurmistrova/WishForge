package model

type Progress struct {
	Progress          float64 `json:"progress"`
	CountCompleted    int     `json:"count_completed"`
	CountNotCompleted int     `json:"count_not_completed"`
}
