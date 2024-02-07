package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
	"strconv"
)

func ItemGet(idUser *int, id *int, idCollection *int) []entities.Item {
	var results *sql.Rows = models.ItemGet(idUser, id, idCollection)
	var items []entities.Item
	for results.Next() {
		var item entities.Item
		err := results.Scan(
			&item.Id,
			&item.IdUser,
			&item.IdItemType,
			&item.IdCollection,
			&item.Title,
			&item.CollectionName,
			&item.IdIgdb,
			&item.Cover,
			&item.Author,
			&item.Year,
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
	cover := ""
	if (item.IdIgdb > 0) {
		cover = IgdbGetCover(strconv.Itoa(item.IdIgdb))
	}

	return models.ItemCreate(
		item.IdUser,
		item.IdItemType,
		item.IdCollection,
		item.Title,
		item.IdIgdb,
		cover,
		item.Author,
		item.Year,
	)
}

func ItemUpdate(item entities.Item) string {
	return models.ItemUpdate(
		item.Id,
		item.IdUser,
		item.Title,
		item.Author,
		item.Year,
	)
}

func ItemDelete(item entities.Item) string {
	return models.ItemDelete(item.Id, item.IdUser)
}
