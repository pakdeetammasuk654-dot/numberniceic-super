package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	numberRepo := repository.NewNumberRepository(db)
	numberHandler := handlers.NewNumberHandler(numberRepo)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/numbers", numberHandler.GetAllNumbers)

}
