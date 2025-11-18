package handlers

import (
	"numberniceic/models"
	"numberniceic/services" // 👈 (อันนี้ยังชี้ไปที่ services)

	"github.com/gofiber/fiber/v2"
)

// 🚀 [เปลี่ยน] เปลี่ยน struct
type AnalysisHandler struct {
	Service services.AnalysisService // 👈 [เปลี่ยน] เปลี่ยน Interface
}

// 🚀 [เปลี่ยน] เปลี่ยนชื่อฟังก์ชัน New
func NewAnalysisHandler(service services.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{Service: service}
}

// (ฟังก์ชัน CalculateAstrology เหมือนเดิม)
func (h *AnalysisHandler) CalculateAstrology(c *fiber.Ctx) error {
	var requestBody models.AstrologyRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if requestBody.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name field is required",
		})
	}

	result, err := h.Service.CalculateNameAstrology(requestBody.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to calculate astrology",
		})
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(result)
}

// (ฟังก์ชัน GetAllSatNums เหมือนเดิม)
// (แม้ชื่อ Handler จะเปลี่ยน แต่ฟังก์ชันนี้ยังทำงานได้)
func (h *AnalysisHandler) GetAllSatNums(c *fiber.Ctx) error {
	results, err := h.Service.GetAllSatNums()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sat_nums",
		})
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(results)
}
