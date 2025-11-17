package handlers

import (
	"numberniceic/models"

	"github.com/gofiber/fiber/v2"
)

type NumberRepository interface {
	GetAllNumbers() ([]models.Number, error)
}

type NumberHandler struct {
	Repo NumberRepository
}

func NewNumberHandler(repo NumberRepository) *NumberHandler {
	return &NumberHandler{Repo: repo}
}

func (h *NumberHandler) GetAllNumbers(c *fiber.Ctx) error {

	results, err := h.Repo.GetAllNumbers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve numbers",
		})
	}

	return c.JSON(results)
}
