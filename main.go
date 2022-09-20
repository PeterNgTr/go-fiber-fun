package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	book "github.com/peterngtr/go-fiber-fun/controller"
	"github.com/peterngtr/go-fiber-fun/database"
	_ "github.com/peterngtr/go-fiber-fun/docs"
	"github.com/peterngtr/go-fiber-fun/model"
)

const PORT string = "3000"
const API_PATH = "/api/v1"
const BOOK_PATH string = "/book"
const DEFAULT_DB = "books.db"

func setupRoutes(app *fiber.App) {
	app.Get(API_PATH+BOOK_PATH, book.GetBooks)

	app.Get(API_PATH+BOOK_PATH+"/:id", book.GetBook)
	app.Post(API_PATH+BOOK_PATH, book.NewBook)
	app.Delete(API_PATH+BOOK_PATH+"/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", DEFAULT_DB)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	faker := faker.New()

	for i := 1; i < 5; i++ {
		var book model.Book
		book.Title = faker.Person().Name()
		book.Author = faker.Person().Name()
		book.Rating = 5
		fmt.Println(book)
		database.DBConn.Create(&book)
	}

	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// @title Get Books APIs
	// @version 1.0
	// @description
	// @host localhost:3000
	// @BasePath /
	app.Get("/docs/*", swagger.HandlerDefault) // default
	initDatabase()

	setupRoutes(app)
	fmt.Printf("Listening at http://localhost:" + PORT)
	app.Listen(":" + PORT)

	defer database.DBConn.Close()
}
