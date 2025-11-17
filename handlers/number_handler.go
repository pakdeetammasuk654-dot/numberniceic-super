package handlers

import (
	"numberniceic/repository"

	"github.com/gofiber/fiber/v2"
)

type NumberHandler struct {
	Repo repository.NumberRepository
}

func NewNumberHandler(repo repository.NumberRepository) *NumberHandler {
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
