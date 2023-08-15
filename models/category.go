package models

// Category représente une catégorie d'ebook.
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
