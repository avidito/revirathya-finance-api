package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpCycleHandler struct {
	cycleUsecase domain.CycleUsecase
}

func NewHttpCycleHandler(app *fiber.App, u domain.CycleUsecase) {
	h := &HttpCycleHandler{
		cycleUsecase: u,
	}

	routes := app.Group("/cycles")
	routes.Post("/", h.CreateCycle)
	routes.Get("/", h.FetchCycles)
	routes.Get("/:id", h.GetCycleByID)
	routes.Put("/:id", h.UpdateCycle)
	routes.Delete("/:id", h.DeleteCycle)
}

// Handler
func (h *HttpCycleHandler) CreateCycle(c *fiber.Ctx) error {
	var err error

	body := CycleRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cycle := domain.Cycle{
		Label:     body.Label,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
		Notes:     body.Notes,
	}

	var createdCycle domain.Cycle
	createdCycle, err = h.cycleUsecase.Create(cycle)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cycleResponse := CycleResponse{
		ID:        createdCycle.ID,
		Label:     createdCycle.Label,
		StartDate: createdCycle.StartDate,
		EndDate:   createdCycle.EndDate,
		Notes:     createdCycle.Notes,
		TotalDays: createdCycle.TotalDays,
	}
	return c.Status(fiber.StatusCreated).JSON(&cycleResponse)
}

func (h *HttpCycleHandler) FetchCycles(c *fiber.Ctx) error {
	var label string
	var err error

	q := c.Queries()
	label = q["label"]

	var cycleReadList []domain.CycleRead
	cycleReadList, err = h.cycleUsecase.Fetch(label)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var cycleReadListResponse []CycleReadResponse
	for _, cycleRead := range cycleReadList {
		cycleReadListResponse = append(
			cycleReadListResponse,
			CycleReadResponse{
				ID:        cycleRead.ID,
				Label:     cycleRead.Label,
				StartDate: cycleRead.StartDate,
				EndDate:   cycleRead.EndDate,
				Notes:     cycleRead.Notes,
				TotalDays: cycleRead.TotalDays,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&cycleReadListResponse)
}

func (h *HttpCycleHandler) GetCycleByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var cycleRead domain.CycleRead
	cycleRead, err = h.cycleUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cycleReadResponse := CycleReadResponse{
		ID:        cycleRead.ID,
		Label:     cycleRead.Label,
		StartDate: cycleRead.StartDate,
		EndDate:   cycleRead.EndDate,
		Notes:     cycleRead.Notes,
		TotalDays: cycleRead.TotalDays,
	}
	return c.Status(fiber.StatusOK).JSON(&cycleReadResponse)
}

func (h *HttpCycleHandler) UpdateCycle(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := CycleRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cycle := domain.Cycle{
		Label:     body.Label,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
		Notes:     body.Notes,
	}

	var updatedCycle domain.Cycle
	updatedCycle, err = h.cycleUsecase.Update(id, cycle)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cycleResponse := CycleResponse{
		ID:        updatedCycle.ID,
		Label:     updatedCycle.Label,
		StartDate: updatedCycle.StartDate,
		EndDate:   updatedCycle.EndDate,
		Notes:     updatedCycle.Notes,
		TotalDays: updatedCycle.TotalDays,
	}
	return c.Status(fiber.StatusOK).JSON(&cycleResponse)
}

func (h *HttpCycleHandler) DeleteCycle(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var deletedCycle domain.Cycle
	deletedCycle, err = h.cycleUsecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cycleResponse := CycleResponse{
		ID:        deletedCycle.ID,
		Label:     deletedCycle.Label,
		StartDate: deletedCycle.StartDate,
		EndDate:   deletedCycle.EndDate,
		Notes:     deletedCycle.Notes,
		TotalDays: deletedCycle.TotalDays,
	}
	return c.Status(fiber.StatusOK).JSON(&cycleResponse)
}
