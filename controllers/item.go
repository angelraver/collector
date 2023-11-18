package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
	"fmt"
)

func ItemGet(id *int, idItemType *int) []entities.Item {
	var results *sql.Rows = models.ItemGet(id, idItemType)
	var items []entities.Item
	fmt.Println(results)

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
			fmt.Println(err)
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
