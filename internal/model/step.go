package model

type Step struct {
	IDWish    int    `json:"id_wish"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}

type NewStep struct {
	IDWish int    `json:"id_wish"`
	Title  string `json:"title"`
}
