package handlers

import (
	"numberniceic/models"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// (Struct และ NewAnalysisHandler เหมือนเดิม)
type AnalysisHandler struct {
	Service services.AnalysisService
}

func NewAnalysisHandler(service services.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{Service: service}
}

// 🚀 [อัปเดต] CalculateAstrology
func (h *AnalysisHandler) CalculateAstrology(c *fiber.Ctx) error {
	var requestBody models.AstrologyRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	// 🚀 [อัปเดต] เพิ่มการตรวจสอบ Day
	if requestBody.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name field is required",
		})
	}
	if requestBody.Day == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Day field is required",
		})
	}

	// 🚀 [อัปเดต] ส่ง Day เข้าไปใน Service
	result, err := h.Service.CalculateNameAstrology(requestBody.Name, requestBody.Day)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to calculate astrology",
		})
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(result)
}

// (GetAllSatNums เหมือนเดิม)
func (h *AnalysisHandler) GetAllSatNums(c *fiber.Ctx) error {
	// ... (โค้ดเดิม) ...
	results, err := h.Service.GetAllSatNums()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sat_nums",
		})
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(results)
}
