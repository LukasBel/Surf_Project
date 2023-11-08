package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"main/models"
	"main/storage"
	"net/http"
	"os"
)

type Repository struct {
	DB *gorm.DB
}

type SurfSpot struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	WaveHeight string `json:"waveHeight"`
	WavePower  int    `json:"wavePower"`
	SkillLevel string `json:"skillLevel"`
}

func (r *Repository) CreateSpot(c *fiber.Ctx) error {
	spotModel := SurfSpot{}
	err := c.BodyParser(&spotModel)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "something went wrong"})
		return err
	}

	err = r.DB.Create(spotModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to create database entry"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "spot created successfully!"})
	return nil

}

func (r *Repository) GetSpots(c *fiber.Ctx) error {
	spotModels := &[]models.SurfSpots{}
	err := r.DB.Find(&spotModels).Error
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "failed to get surf spots"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "spots found successfully!", "data": spotModels})
	return nil
}

func (r *Repository) DeleteSpot(c *fiber.Ctx) error {
	id := c.Params("id")

	surfModel := &models.SurfSpots{}

	if id == "" {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "empty ID"})
		return nil
	}

	err := r.DB.Where("id = ?", id).Delete(&surfModel).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to delete spot"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "spot deleted successfully"})
	return nil

}

func (r *Repository) SetUpRoutes(App *fiber.App) {
	api := App.Group("/api")
	api.Get("/spots", r.GetSpots)
	api.Post("/create_spot", r.CreateSpot)
	api.Delete("/delete/:id", r.DeleteSpot)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = models.MigrateSpots(db)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetUpRoutes(app)
	app.Listen(":8080")
}
