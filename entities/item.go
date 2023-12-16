package entities

import "time"

type Item struct {
	Id         int       `json:"id"`
	IdUser     int       `json:"iduser"`
	IdItemType int       `json:"iditemtype"`
	IdCollection int       `json:"idcollection"`
	Title      string    `json:"title"`
	CollectionName   string    `json:"collectionname"`
	IdIgdb     int       `json:"idigdb"`
	Cover			string 			`json:"cover"`
	CreatedAt  time.Time `json:"createdat"`
	UpdatedAt  time.Time `json:"updatedat"`
}
