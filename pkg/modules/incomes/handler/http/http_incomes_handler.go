package http

import (
	"strconv"
	"time"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type ResponseError struct {
	Message string `json:"message"`
}

type HttpIncomeHandler struct {
	incomeUsecase domain.IncomeUsecase
}

func NewHttpIncomeHandler(app *fiber.App, u domain.IncomeUsecase) {
	h := &HttpIncomeHandler{
		incomeUsecase: u,
	}

	routes := app.Group("/incomes")
	routes.Post("/", h.CreateIncome)
	routes.Get("/", h.FetchIncomes)
	routes.Get("/:id", h.GetIncome)
	routes.Put("/:id", h.UpdateIncome)
	routes.Delete("/:id", h.DeleteIncome)
}

// Handler
func (h *HttpIncomeHandler) CreateIncome(c *fiber.Ctx) (err error) {
	body := CreateIncomeRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	income := domain.Income{
		Date:         body.Date,
		IncomeTypeID: body.IncomeTypeID,
		LocationID:   body.LocationID,
		Description:  body.Description,
		Amount:       body.Amount,
	}

	income, err = h.incomeUsecase.Create(income)
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	incomeResponse := IncomeResponse{
		ID:           income.ID,
		Date:         income.Date,
		IncomeTypeID: income.IncomeTypeID,
		LocationID:   income.LocationID,
		Description:  income.Description,
		Amount:       income.Amount,
	}

	return c.Status(fiber.StatusCreated).JSON(&incomeResponse)
}

func (h *HttpIncomeHandler) FetchIncomes(c *fiber.Ctx) (err error) {
	var _date domain.DateStandard
	var income_type string

	q := c.Queries()
	_date.Time, err = time.Parse("2006-01-02", q["date"])
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	income_type = q["income_type"]

	var incomes []domain.IncomeRead
	incomes, err = h.incomeUsecase.Fetch(_date, income_type)
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var incomeResponseSubs []IncomeResponseSubs
	var tmpIncomeResponseSubs IncomeResponseSubs
	for i, income := range incomes {
		tmpIncomeResponseSubs = IncomeResponseSubs{
			Date:        income.Date,
			IncomeType:  income.IncomeType,
			Location:    income.Location,
			Description: income.Description,
			Amount:      income.Amount,
		}
		incomeResponseSubs = append(incomeResponseSubs, tmpIncomeResponseSubs)
	}

	return c.Status(fiber.StatusOK).JSON(&incomeResponseSubs)
}

func (h *HttpIncomeHandler) GetIncome(c *fiber.Ctx) (err error) {
	var id int64
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var income domain.IncomeRead
	income, err = h.incomeUsecase.Get(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	incomeResponseSubs := IncomeResponseSubs{
		Date:         income.Date,
		IncomeTypeID: income.IncomeTypeID,
		LocationID:   income.LocationID,
		Description:  income.Description,
		Amount:       income.Amount,
	}

	return c.Status(fiber.StatusOK).JSON(&income)
}

func (h *HttpIncomeHandler) UpdateIncome(c *fiber.Ctx) (err error) {
	var id int64
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	body := UpdateIncomeRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var income domain.Income
	income, err = h.incomeUsecase.Update(id, income)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	incomeResponseSubs := IncomeResponseSubs{
		Date:        body.Date,
		IncomeType:  body.IncomeType,
		Location:    body.Location,
		Description: body.Description,
		Amount:      body.Amount,
	}

	return c.Status(fiber.StatusOK).JSON(&income)
}

func (h *HttpIncomeHandler) DeleteIncome(c *fiber.Ctx) (err error) {
	var id int64
	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var income domain.Income
	income, err = h.incomeUsecase.Delete(id)
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&income)
}
