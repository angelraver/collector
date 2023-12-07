package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
)

func ItemGet(idUser *int, id *int, idItemType *int) []entities.Item {
	var results *sql.Rows = models.ItemGet(idUser, id, idItemType)
	var items []entities.Item

	for results.Next() {
		var item entities.Item
		err := results.Scan(
			&item.Id,
			&item.IdUser,
			&item.IdItemType,
			&item.Title,
			&item.TypeName,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil
		}
		items = append(items, item)
	}
	return items
}

func ItemCreate(item entities.Item) string {
	return models.ItemCreate(item.IdUser, item.IdItemType, item.Title)
}

func ItemUpdate(item entities.Item) string {
	return models.ItemUpdate(item.Id, item.IdItemType, item.Title)
}

func ItemDelete(item entities.Item) string {
	return models.ItemDelete(item.Id, item.IdUser)
}
