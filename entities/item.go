package entities

import "time"

type Item struct {
	Id         int       `json:"id"`
	IdUser     int       `json:"iduser"`
	IdItemType int       `json:"iditemtype"`
	Title      string    `json:"title"`
	TypeName   string    `json:"typename"`
	CreatedAt  time.Time `json:"createdat"`
	UpdatedAt  time.Time `json:"updatedat"`
}
