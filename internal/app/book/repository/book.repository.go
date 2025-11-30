package repository

import (
	"sistem-buku/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(title string, author string, publisher string, publishYear int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(title string, author string, publisher string, publishYear int) error {
	book := entity.Book{
		ID:          uuid.New(),
		Title:       title,
		Author:      author,
		Publisher:   publisher,
		PublishYear: publishYear,
	}
	return r.db.Create(&book).Error
}
