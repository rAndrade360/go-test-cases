package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(name, email string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(name, email string) error {
	log.Println("User created!")
	return nil
}
