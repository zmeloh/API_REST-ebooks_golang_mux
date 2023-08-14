package controllers

import (
	"example/api/dao"
	"example/api/models"
)


func CreateEbook(e models.Ebook){


}


func GetAllEbooks(){


}

func GetEbookByID(id int){


}

func GetEbooksByCategoryID(id int){


}



func UpdateEbook(id int, old models.Ebook) {
	ebook := dao.SelectEbookByID(id)
	ebook.Title = old.Title
	ebook.Author = old.Author
	ebook.CategoryID = old.CategoryID
	dao.UpdateFavorite(ebook)
}


func DeleteEbook(id int){


}
