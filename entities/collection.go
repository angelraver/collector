package entities

import "time"

type Collection struct {
	Id        int       `json:"id"`
	IdUser    int       `json:"iduser"`
	IdItemType      int    `json:"iditemtype"`
	Name      string    `json:"name"`
	ItemsCount      string    `json:"itemscount"`
	IdPlatform      int    `json:"idplatform"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type CollectionBkp struct {
	Id        int       `json:"id"`
	IdUser    int       `json:"iduser"`
	IdItemType      int    `json:"iditemtype"`
	Name      string    `json:"name"`
	IdPlatform      int    `json:"idplatform"`
}
