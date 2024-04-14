package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
)

func CollectionBkp() []entities.CollectionBkp {
	var rows *sql.Rows = models.CollectionGetAll()
	var collections []entities.CollectionBkp

	for rows.Next() {
		var collection entities.CollectionBkp
		err := rows.Scan(
			&collection.Id,
			&collection.IdUser,
			&collection.Name,
			&collection.IdItemType,
			&collection.IdPlatform,
		)
		if err != nil {
			return nil
		}
		collections = append(collections, collection)
	}
	return collections
}

func CollectionCreate(collection entities.Collection) map[string]interface{} {
	row := models.CollectionCreate(
		collection.IdUser,
		collection.IdItemType,
		collection.Name,
		collection.IdPlatform,
	)

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

func CollectionDelete(collection entities.Collection) string {
	items := ItemGet(&collection.IdUser, nil, &collection.Id)
	for _, item := range items {
		ItemDelete(item)
	}

	return models.CollectionDelete(collection.Id, collection.IdUser)
}

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
			&collection.IdPlatform,
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

func CollectionUpdate(collection entities.Collection) string {
	return models.CollectionUpdate(
		collection.Id,
		collection.IdUser,
		collection.Name,
	  collection.IdPlatform,
	)
}