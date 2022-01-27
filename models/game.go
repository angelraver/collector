package models

type Game struct {
	Id        int    `json:"id"`
	IdConsole int    `json:"idConsole"`
	Title     string `json:"title"`
	Stars     int    `json:"stars"`
	Qty       int    `json:"qty"`
}
