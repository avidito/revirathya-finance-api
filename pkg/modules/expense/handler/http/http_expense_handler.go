package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpExpenseHandler struct {
	expenseUsecase domain.ExpenseUsecase
}

func NewHttpExpenseHandler(app *fiber.App, u domain.ExpenseUsecase) {
	h := &HttpExpenseHandler{
		expenseUsecase: u,
	}

	routes := app.Group("/expenses")
	routes.Post("/", h.CreateExpense)
	routes.Get("/", h.FetchExpenses)
	routes.Get("/:id", h.GetExpenseByID)
	routes.Put("/:id", h.UpdateExpense)
	routes.Delete("/:id", h.DeleteExpense)
}

// Handler
func (h *HttpExpenseHandler) CreateExpense(c *fiber.Ctx) error {
	var err error

	body := ExpenseRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	expense := domain.Expense{
		Date:         body.Date,
		BudgetTypeID: body.BudgetTypeID,
		LocationID:   body.LocationID,
		Description:  body.Description,
		Amount:       body.Amount,
	}

	var createdExpense domain.Expense
	createdExpense, err = h.expenseUsecase.Create(expense)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	expenseResponse := ExpenseResponse{
		ID:           createdExpense.ID,
		Date:         createdExpense.Date,
		BudgetTypeID: createdExpense.BudgetTypeID,
		LocationID:   createdExpense.LocationID,
		Description:  createdExpense.Description,
		Amount:       createdExpense.Amount,
	}
	return c.Status(fiber.StatusCreated).JSON(&expenseResponse)
}

func (h *HttpExpenseHandler) FetchExpenses(c *fiber.Ctx) error {
	var _date string
	var budget_type string
	var err error

	q := c.Queries()
	_date = q["date"]
	budget_type = q["budget_type"]

	var expenseReadList []domain.ExpenseRead
	expenseReadList, err = h.expenseUsecase.Fetch(_date, budget_type)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var expenseReadListResponse []ExpenseReadResponse
	for _, expenseRead := range expenseReadList {
		expenseReadListResponse = append(
			expenseReadListResponse,
			ExpenseReadResponse{
				Date:        expenseRead.Date,
				BudgetType:  expenseRead.BudgetType,
				Location:    expenseRead.Location,
				Description: expenseRead.Description,
				Amount:      expenseRead.Amount,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&expenseReadListResponse)
}

func (h *HttpExpenseHandler) GetExpenseByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var expenseRead domain.ExpenseRead
	expenseRead, err = h.expenseUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	expenseReadResponse := ExpenseReadResponse{
		Date:        expenseRead.Date,
		BudgetType:  expenseRead.BudgetType,
		Location:    expenseRead.Location,
		Description: expenseRead.Description,
		Amount:      expenseRead.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&expenseReadResponse)
}

func (h *HttpExpenseHandler) UpdateExpense(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := ExpenseRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	expense := domain.Expense{
		Date:         body.Date,
		BudgetTypeID: body.BudgetTypeID,
		LocationID:   body.LocationID,
		Description:  body.Description,
		Amount:       body.Amount,
	}

	var updatedExpense domain.Expense
	updatedExpense, err = h.expenseUsecase.Update(id, expense)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	expenseResponse := ExpenseResponse{
		ID:           updatedExpense.ID,
		Date:         updatedExpense.Date,
		BudgetTypeID: updatedExpense.BudgetTypeID,
		LocationID:   updatedExpense.LocationID,
		Description:  updatedExpense.Description,
		Amount:       updatedExpense.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&expenseResponse)
}

func (h *HttpExpenseHandler) DeleteExpense(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var deletedExpense domain.Expense
	deletedExpense, err = h.expenseUsecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	expenseResponse := ExpenseResponse{
		ID:           deletedExpense.ID,
		Date:         deletedExpense.Date,
		BudgetTypeID: deletedExpense.BudgetTypeID,
		LocationID:   deletedExpense.LocationID,
		Description:  deletedExpense.Description,
		Amount:       deletedExpense.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&expenseResponse)
}
