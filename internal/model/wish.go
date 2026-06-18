package model

type Wish struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}

type NewWish struct {
	Title string `json:"title"`
}
