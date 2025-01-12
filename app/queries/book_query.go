package queries

import (
	"gorm-authentication/app/model"
)

func GetBook(id uint) (*model.Book, error) {
	var book model.Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func GetBooks() ([]model.Book, error) {
	var books []model.Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func CreateBook(book *model.Book) error {
	result := db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateBookBook(book *model.Book) error {
	result := db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
