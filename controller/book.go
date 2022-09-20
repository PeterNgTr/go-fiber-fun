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

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return c.JSON(books)
}

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

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(NOT_FOUND_ID)
	}
	db.Create(&book)
	return c.JSON(book)
}

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
