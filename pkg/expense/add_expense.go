package expense

import (
	"github.com/avidito/revirathya-finance-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddExpenseRequestBody struct {
	BudgetType  string `json:"budget_type"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func (h handler) AddExpense(c *fiber.Ctx) error {
	body := AddExpenseRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var expense models.Expense
	expense.BudgetType = body.BudgetType
	expense.Location = body.Location
	expense.Description = body.Description
	expense.Amount = body.Amount

	if result := h.DB.Create(&expense); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&expense)
}
