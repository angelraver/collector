package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
)

func ItemTypeGet(idUser *int, id *int) []entities.ItemType {
	var results *sql.Rows = models.ItemTypeGet(idUser, id)
	var items []entities.ItemType

	for results.Next() {
		var itemType entities.ItemType
		err := results.Scan(
			&itemType.Id,
			&itemType.IdUser,
			&itemType.Name,
			&itemType.CreatedAt,
			&itemType.UpdatedAt,
		)
		if err != nil {
			return nil
		}
		items = append(items, itemType)
	}
	return items
}

func ItemTypeCreate(itemType entities.ItemType) string {
	return models.ItemTypeCreate(itemType.IdUser, itemType.Name)
}

func ItemTypeUpdate(itemType entities.ItemType) string {
	return models.ItemTypeUpdate(itemType.Id, itemType.Name)
}
