package usecase

import (
	"sistem-buku/internal/app/book/repository"
	"sistem-buku/internal/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type BookUsecase interface {
	CreateBook(entity.Book) error
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) CreateBook(book entity.Book) error {
	if book.PublishYear < 2020 {
		return fiber.NewError(fiber.StatusBadRequest, "Publish year must be 2020 or later")
	}

	return u.bookRepo.CreateBook(book.Title, book.Author, book.Publisher, book.PublishYear)
}
