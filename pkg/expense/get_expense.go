package expense

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetExpense(c *fiber.Ctx) error {
	id := c.Params("id")

	var expense models.Expense
	if result := h.DB.First(&expense, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&expense)
}
