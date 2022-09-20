package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/peterngtr/go-fiber-fun/database"
	"github.com/peterngtr/go-fiber-fun/model"
)

const NOT_FOUND_ID string = "No Book Found with ID"
const SUCCESSFUL_DELETION string = "Book Successfully deleted"

// GetBooks example
// @Tags Books
// @Summary Get all books
// @Description Get all books
// @Produce  json
// @Success 200 {array} model.Book "ok"
// @Router /api/v1/book [get]
func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return c.JSON(books)
}

// GetBookWithId example
// @Tags Books
// @Summary Get book by ID
// @Description Get book by Id
// @Param        id   path      int  true  "Book ID"
// @Produce  json
// @Success 200 {object} model.Book "ok"
// @Router /api/v1/book/{id} [get]
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	err := db.Find(&book, id).Error
	fmt.Println(db.Find(&book, id).Error)
	if err != nil {
		return c.Status(500).SendString(NOT_FOUND_ID)
	}
	return c.JSON(book)
}

// @Tags Books
// @Summary Create new book
// @Description Create new book
// @Produce  json
// @Param        book  body      model.Book  true  "Add book"
// @Success 200 {object} model.Book "ok"
// @Router /api/v1/book [post]
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(NOT_FOUND_ID)
	}
	db.Create(&book)
	return c.JSON(book)
}

// @Tags Books
// @Summary Delete book by id
// @Description Delete book by id
// @Produce  json
// @Param        id   path      int  true  "Book ID"
// @Success 200 {string} string "Book Successfully deleted"
// @Router /api/v1/book/{id} [delete]
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	err := db.First(&book, id).Error
	if err != nil {
		return c.Status(500).SendString(NOT_FOUND_ID)
	}
	db.Delete(&book)
	return c.SendString(SUCCESSFUL_DELETION)
}
