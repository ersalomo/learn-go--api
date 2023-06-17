package service

import (
	"go-api/db/model"
	"go-api/db/repository"
)

type BookService interface {
	FindAll() ([]model.Book, error)
	FindById(id int) (model.Book, error)
	Create(book model.Book) (model.Book, error)
}

type booksService struct {
	repo repository.BookRepo
}

func NewBookService(repo repository.BookRepo) *booksService {
	return &booksService{repo}
}

func (bs *booksService) FindAll() ([]model.Book, error) {
	return bs.repo.FindAll()
}
func (bs *booksService) FindById(id int) (model.Book, error) {
	return bs.repo.FindById(id)
}

func (bs *booksService) Create(book model.Book) (model.Book, error) {
	return bs.repo.Create(book)

}
