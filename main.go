package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"main/storage"
	"net/http"
)

type Repository struct {
	DB gorm.DB
}

type SurfSpot struct {
	Type       string `json:"type"`
	WaveHeight string `json:"waveHeight"`
	WavePower  int    `json:"wavePower"`
	SkillLevel string `json:"skillLevel"`
}

func (r *Repository) CreateSpot(c *fiber.Ctx) error {
	spotModel := SurfSpot{}
	err := c.BodyParser(&spotModel)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "something went wrong"})
	}

	err = r.DB.Create(spotModel).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "failed to create database entry"})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "spot created successfully!"})
	return nil

}

func (r *Repository) GetSpots(c *fiber.Ctx) error {
	spotModels := &[]storage.SurfSpots{}
	err := r.DB.Find(spotModels)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "failed to get surf spots"})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "spots found successfully!", "data": spotModels})
	return nil
}

func (r *Repository) SetUpRoutes(App *fiber.App) {
	api := App.Group("/api")
	api.Get("/spots", r.GetSpots)
	api.Post("/create_spot", r.CreateSpot)
	api.Delete("/delete/:id", r.DeleteSpot)
}

func main() {
	r := Repository{
		DB: db,
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	app := fiber.New()
	r.SetUpRoutes(app)

	app.Listen(":8080")
}
