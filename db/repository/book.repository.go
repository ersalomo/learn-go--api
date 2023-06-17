package repository

import (
	"go-api/db/model"
	"gorm.io/gorm"
)

type BookRepo interface {
	FindAll() ([]model.Book, error)
	FindById(id int) (model.Book, error)
	Create(book model.Book) (model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := br.db.Find(&books)
	return books, err.Error

}
func (br *bookRepository) FindById(id int) (model.Book, error) {
	var book model.Book
	err := br.db.Where("id = ?", id).First(&book)
	return book, err.Error
}

func (br *bookRepository) Create(book model.Book) (model.Book, error) {
	err := br.db.Create(&book)
	return book, err.Error
}
