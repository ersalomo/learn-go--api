package repository

import (
	"go-api/db/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	FindByEmail(email string) (model.User, error)
	Create(user model.User) (model.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users)
	return users, err.Error
}
func (r *repository) FindById(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ? ", id).First(&user)
	//err := r.db.Find(&user, id)
	return user, err.Error
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user)
	return user, err.Error
}

func (r *repository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user)
	return user, err.Error
}
