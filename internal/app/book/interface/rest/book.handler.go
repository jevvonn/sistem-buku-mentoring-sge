package rest

import (
	"sistem-buku/internal/app/book/usecase"
	"sistem-buku/internal/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(app *fiber.App, bookUsecase usecase.BookUsecase) {
	handler := &BookHandler{
		bookUsecase: bookUsecase,
	}

	app.Post("/books", handler.CreateBook)
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	type request struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PublishYear int    `json:"publish_year"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	book := entity.Book{
		Title:       req.Title,
		Author:      req.Author,
		Publisher:   req.Publisher,
		PublishYear: req.PublishYear,
	}
	if err := h.bookUsecase.CreateBook(book); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
	})
}
