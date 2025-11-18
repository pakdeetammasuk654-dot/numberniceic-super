package handlers

import (
	"errors"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

type NumberHandler struct {
	Service services.NumberService // 👈 เปลี่ยนจาก Repo เป็น Service
}

func NewNumberHandler(service services.NumberService) *NumberHandler {
	return &NumberHandler{Service: service}
}

func (h *NumberHandler) GetNumberByPairNumber(c *fiber.Ctx) error {

	pairNumber := c.Params("number")

	number, err := h.Service.GetNumberByPairNumber(pairNumber)
	if err != nil {
		// 4. [สำคัญ] แยกแยะประเภท Error
		if errors.Is(err, repository.ErrNotFound) {
			// ถ้า "ไม่พบ" (Service ส่งต่อมาจาก Repo)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Number not found",
			})
		}

		// ถ้าเป็น "Input ผิด" (Service สร้างเอง)
		if err.Error() == "pairnumber must be 2 characters" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// ถ้าเป็น Error ร้ายแรงอื่นๆ
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(number)
}

func (h *NumberHandler) GetAllNumbers(c *fiber.Ctx) error {

	results, err := h.Service.GetAllNumbers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve numbers",
		})
	}

	return c.JSON(results)
}
