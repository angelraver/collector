package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
)

func ItemTypeGet(idUser *int, id *int) []entities.ItemType {
	var rows *sql.Rows = models.ItemTypeGet(idUser, id)
	var items []entities.ItemType

	for rows.Next() {
		var itemType entities.ItemType
		err := rows.Scan(
			&itemType.Id,
			&itemType.IdUser,
			&itemType.Name,
			&itemType.ItemsCount,
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

func ItemTypeCreate(itemType entities.ItemType)  map[string]interface{} {
	row := models.ItemTypeCreate(itemType.IdUser, itemType.Name)

	var newID int
	err := row.Scan(&newID)
	if err != nil {
		return nil
	}

	result := map[string]interface{}{
		"message": "Collection "+itemType.Name+" created!", 
		"id": newID,
	}

	return result
}

func ItemTypeUpdate(itemType entities.ItemType) string {
	return models.ItemTypeUpdate(itemType.Id, itemType.Name)
}

func ItemTypeDelete(itemType entities.ItemType) string {
	return models.ItemTypeDelete(itemType.Id, itemType.IdUser)
}
