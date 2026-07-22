package model

type Step struct {
	IDWish    int    `json:"id_wish"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}
type CreateStep struct {
	IDWish int    `json:"id_wish"`
	Title  string `json:"title"`
}

type CreateStepRequest struct {
	Title string `json:"title"`
}
type UpdateStepRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}
