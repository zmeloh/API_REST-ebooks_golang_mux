package services

import (
	"example/api/controllers"
	"example/api/models"
)

func UpdateEbook(id int, e models.Ebook) {
	controllers.UpdateEbook(id, e)
	return
}


