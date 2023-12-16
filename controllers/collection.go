package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
)

func CollectionGet(idUser *int, id *int) []entities.Collection {
	var rows *sql.Rows = models.CollectionGet(idUser, id)
	var collections []entities.Collection

	for rows.Next() {
		var collection entities.Collection
		err := rows.Scan(
			&collection.Id,
			&collection.IdUser,
			&collection.Name,
			&collection.ItemsCount,
			&collection.IdItemType,
			&collection.CreatedAt,
			&collection.UpdatedAt,
		)
		if err != nil {
			return nil
		}
		collections = append(collections, collection)
	}
	return collections
}

func CollectionCreate(collection entities.Collection) map[string]interface{} {
	row := models.CollectionCreate(collection.IdUser, collection.IdItemType, collection.Name)

	var newID int
	err := row.Scan(&newID)
	if err != nil {
		return nil
	}

	result := map[string]interface{}{
		"message": "Collection " + collection.Name + " created!", 
		"id": newID,
	}

	return result
}

func CollectionUpdate(collection entities.Collection) string {
	return models.CollectionUpdate(collection.Id, collection.IdUser, collection.Name)
}

func CollectionDelete(collection entities.Collection) string {
	return models.CollectionDelete(collection.Id, collection.IdUser)
}
