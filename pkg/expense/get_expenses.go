package expense

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	if result := h.DB.Find(&expenses); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&expenses)
}
