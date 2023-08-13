package models

type Favorite struct{
	ID int `json:"id"`
	UserID int `json:"user_id"`
	EbookID int `json:"ebook_id"`
}