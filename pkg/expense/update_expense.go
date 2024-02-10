package expense

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateExpenseRequestBody struct {
	BudgetType  string `json:"budget_type"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func (h handler) UpdateExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateExpenseRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var expense models.Expense
	if result := h.DB.First(&expense, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	expense.BudgetType = body.BudgetType
	expense.Location = body.Location
	expense.Description = body.Description
	expense.Amount = body.Amount

	h.DB.Save(expense)
	return c.Status(fiber.StatusOK).JSON(&expense)
}
