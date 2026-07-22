package model

type Wish struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}

type CreateWishRequest struct {
	Title string `json:"title"`
}
type UpdateWishRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"is_completed"`
}
