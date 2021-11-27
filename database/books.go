package database 

import (
	"github.com/jinzhu/gorm"
	"github.com/felipeas314/go/models"
)


func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*").
			Group("books.id")
	if err := query.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}