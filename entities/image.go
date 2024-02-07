package entities

import "time"

type Image struct {
	Id         int       `json:"id"`
	IdUser     int       `json:"iduser"`
	IdItem     int       `json:"iditem"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdat"`
	UpdatedAt  time.Time `json:"updatedat"`
}
