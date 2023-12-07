package entities

import "time"

type ItemType struct {
	Id        int       `json:"id"`
	IdUser    int       `json:"iduser"`
	Name      string    `json:"name"`
	ItemsCount      string    `json:"itemscount"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
