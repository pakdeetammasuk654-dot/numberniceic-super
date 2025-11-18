package handlers

import (
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// SatNumHandler handles the HTTP requests for SatNum.
type SatNumHandler struct {
	Service services.SatNumService
}

// NewSatNumHandler creates a new SatNumHandler.
func NewSatNumHandler(service services.SatNumService) *SatNumHandler {
	return &SatNumHandler{Service: service}
}

// GetAllSatNums is the handler for GET /api/v1/satnums
func (h *SatNumHandler) GetAllSatNums(c *fiber.Ctx) error {
	results, err := h.Service.GetAllSatNums()
	if err != nil {
		// ถ้ามี error, ส่ง 500 Internal Server Error กลับไป
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sat_nums",
		})
	}

	// ส่งผลลัพธ์กลับไปเป็น JSON
	return c.JSON(results)
}
