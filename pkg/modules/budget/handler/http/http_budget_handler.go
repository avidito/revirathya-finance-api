package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpBudgetHandler struct {
	budgetUsecase domain.BudgetUsecase
}

func NewHttpBudgetHandler(app *fiber.App, u domain.BudgetUsecase) {
	h := &HttpBudgetHandler{
		budgetUsecase: u,
	}

	routes := app.Group("/budgets")
	routes.Post("/", h.CreateBudget)
	routes.Get("/", h.FetchBudgets)
	routes.Get("/:id", h.GetBudgetByID)
	routes.Put("/:id", h.UpdateBudget)
	routes.Delete("/:id", h.DeleteBudget)
}

// Handler
func (h *HttpBudgetHandler) CreateBudget(c *fiber.Ctx) error {
	var err error

	body := BudgetRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	budget := domain.Budget{
		Cycle:        body.Cycle,
		BudgetTypeID: body.BudgetTypeID,
		Amount:       body.Amount,
	}

	var createdBudget domain.Budget
	createdBudget, err = h.budgetUsecase.Create(budget)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	budgetResponse := BudgetResponse{
		ID:           createdBudget.ID,
		Cycle:        createdBudget.Cycle,
		BudgetTypeID: createdBudget.BudgetTypeID,
		Amount:       createdBudget.Amount,
	}
	return c.Status(fiber.StatusCreated).JSON(&budgetResponse)
}

func (h *HttpBudgetHandler) FetchBudgets(c *fiber.Ctx) error {
	var cycle string
	var budget_type string
	q := c.Queries()
	cycle = q["cycle"]
	budget_type = q["budget_type"]

	var budgetReadList []domain.BudgetRead
	var err error
	budgetReadList, err = h.budgetUsecase.Fetch(cycle, budget_type)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var budgetReadListResponse []BudgetReadResponse
	for _, budgetRead := range budgetReadList {
		budgetReadListResponse = append(
			budgetReadListResponse,
			BudgetReadResponse{
				Cycle:      budgetRead.Cycle,
				BudgetType: budgetRead.BudgetType,
				Amount:     budgetRead.Amount,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&budgetReadListResponse)
}

func (h *HttpBudgetHandler) GetBudgetByID(c *fiber.Ctx) error {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var budgetRead domain.BudgetRead
	budgetRead, err = h.budgetUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	budgetReadResponse := BudgetReadResponse{
		Cycle:      budgetRead.Cycle,
		BudgetType: budgetRead.BudgetType,
		Amount:     budgetRead.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&budgetReadResponse)
}

func (h *HttpBudgetHandler) UpdateBudget(c *fiber.Ctx) error {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := BudgetRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	budget := domain.Budget{
		Cycle:        body.Cycle,
		BudgetTypeID: body.BudgetTypeID,
		Amount:       body.Amount,
	}

	var updatedBudget domain.Budget
	updatedBudget, err = h.budgetUsecase.Update(id, budget)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	budgetResponse := BudgetResponse{
		ID:           updatedBudget.ID,
		Cycle:        updatedBudget.Cycle,
		BudgetTypeID: updatedBudget.BudgetTypeID,
		Amount:       updatedBudget.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&budgetResponse)
}

func (h *HttpBudgetHandler) DeleteBudget(c *fiber.Ctx) error {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var deletedBudget domain.Budget
	deletedBudget, err = h.budgetUsecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	budgetResponse := BudgetResponse{
		ID:           deletedBudget.ID,
		Cycle:        deletedBudget.Cycle,
		BudgetTypeID: deletedBudget.BudgetTypeID,
		Amount:       deletedBudget.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&budgetResponse)
}
