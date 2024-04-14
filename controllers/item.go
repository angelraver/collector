package controllers

import (
	"coleccionista/entities"
	"coleccionista/models"
	"database/sql"
	"strconv"
)

func ItemCreate(item entities.Item) int {
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

func ItemBkp() []entities.ItemBkp {
	var results *sql.Rows = models.ItemGetAll()
	var items []entities.ItemBkp
	for results.Next() {
		var item entities.ItemBkp
		err := results.Scan(
			&item.Id,
			&item.IdUser,
			&item.IdItemType,
			&item.IdCollection,
			&item.Title,
			&item.IdIgdb,
			&item.Cover,
			&item.Author,
			&item.Year,
		)
		if err != nil {
			return nil
		}
		items = append(items, item)
	}
	return items
}

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

func ItemDelete(item entities.Item) string {
	images := ImageGet(&item.IdUser, &item.Id)
	for _, image := range images {
		ImageDelete(image)
	}
	return models.ItemDelete(item.Id, item.IdUser)
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
