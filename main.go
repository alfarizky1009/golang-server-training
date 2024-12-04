package main

import (
	"fmt"
	"golang-server-training-postgres/models"
	"golang-server-training-postgres/storage"
	"log"
	"os"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Interface for DB
type Repository struct {
	DB *gorm.DB
}

// fiber is system working on background to do the request and response
// We able to get response body from fiber context
func (r *Repository) CreateBooks(context *fiber.Ctx) error {
	book := models.Books{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Failed to create the book"})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book has been added"})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Failed to get the book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Books fetched successfully",
			"data": bookModels,
		})

	return nil
}

func (r *Repository) DeleteBooks(context *fiber.Ctx)  error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "ID Cannot be empty"})

		return nil
	}

	err := r.DB.Delete(bookModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete book"})

			return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "books has been deleted successfully"})

	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Id cannot be empty"})
		return nil
	}

	fmt.Println("the ID is ", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get the book"})
			return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book id fetched successfully",
			"data": bookModel,
	})

	return nil
}



// Struct method
func (r *Repository) SetupRoutes(app *fiber.App){
	// All our API will start with /api
	api := app.Group("/api")

	api.Post("/create_books", r.CreateBooks)
	api.Delete("/delete_books/:id", r.DeleteBooks)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}

func main() {
	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		Database: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load Database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	// Create repository
	r:= Repository {
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8761")
}