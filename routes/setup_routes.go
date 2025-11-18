package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {

	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)

}
