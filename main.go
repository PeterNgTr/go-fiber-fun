package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/peterngtr/go-fiber-fun/config"
	book "github.com/peterngtr/go-fiber-fun/controller"
	"github.com/peterngtr/go-fiber-fun/database"
	_ "github.com/peterngtr/go-fiber-fun/docs"
	"github.com/peterngtr/go-fiber-fun/middleware"
	"github.com/peterngtr/go-fiber-fun/model"
	"golang.org/x/crypto/bcrypt"
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
	database.DBConn.AutoMigrate(model.Book{}, model.User{})
	faker := faker.New()

	for i := 1; i < 5; i++ {
		var book model.Book
		book.Title = faker.Person().Name()
		book.Author = faker.Person().Name()
		book.Rating = 5
		database.DBConn.Create(&book)
	}

	var user model.User
	hash, err := hashPassword(config.Config("DEFAULT_PASS"))
	user.Email = "helloworld@test.de"
	user.Password = hash
	user.Username = "admin"
	database.DBConn.Create(&user)
	fmt.Println("Database Migrated")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// Login route
	app.Post("/login", middleware.Login)

	// @title Get Books APIs
	// @version 1.0
	// @description
	// @host localhost:3000
	// @BasePath /
	app.Get("/docs/*", swagger.HandlerDefault) // default
	initDatabase()

	// JWT Middleware
	//app.Use(jwtware.New(jwtware.Config{
	//SigningKey: []byte(os.Getenv("SECRET")),
	//}))

	setupRoutes(app)
	fmt.Printf("Listening at http://localhost:" + PORT)
	app.Listen(":" + PORT)

	defer database.DBConn.Close()
}
