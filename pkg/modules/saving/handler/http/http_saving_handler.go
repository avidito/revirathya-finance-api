package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpSavingHandler struct {
	tranferUsecase domain.SavingUsecase
}

func NewHttpSavingHandler(app *fiber.App, u domain.SavingUsecase) {
	h := &HttpSavingHandler{
		tranferUsecase: u,
	}

	routes := app.Group("/savings")
	routes.Post("/", h.CreateSaving)
	routes.Get("/", h.FetchSavings)
	routes.Get("/:id", h.GetSavingByID)
	routes.Put("/:id", h.UpdateSaving)
	routes.Delete("/:id", h.DeleteSaving)
}

// Handler
func (h *HttpSavingHandler) CreateSaving(c *fiber.Ctx) error {
	var err error

	body := SavingRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tranfer := domain.Saving{
		Date:         body.Date,
		SavingTypeID: body.SavingTypeID,
		Flow:         body.Flow,
		Amount:       body.Amount,
	}

	var createdSaving domain.Saving
	createdSaving, err = h.tranferUsecase.Create(tranfer)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := SavingResponse{
		ID:           createdSaving.ID,
		Date:         createdSaving.Date,
		SavingTypeID: createdSaving.SavingTypeID,
		Flow:         createdSaving.Flow,
		Amount:       createdSaving.Amount,
	}
	return c.Status(fiber.StatusCreated).JSON(&tranferResponse)
}

func (h *HttpSavingHandler) FetchSavings(c *fiber.Ctx) error {
	var date string
	var saving_type string
	var flow string
	var err error

	q := c.Queries()
	date = q["date"]
	saving_type = q["saving_type"]
	flow = q["flow"]

	var tranferReadList []domain.SavingRead
	tranferReadList, err = h.tranferUsecase.Fetch(date, saving_type, flow)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var tranferReadListResponse []SavingReadResponse
	for _, tranferRead := range tranferReadList {
		tranferReadListResponse = append(
			tranferReadListResponse,
			SavingReadResponse{
				Date:       tranferRead.Date,
				SavingType: tranferRead.SavingType,
				Flow:       tranferRead.Flow,
				Amount:     tranferRead.Amount,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&tranferReadListResponse)
}

func (h *HttpSavingHandler) GetSavingByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var tranferRead domain.SavingRead
	tranferRead, err = h.tranferUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferReadResponse := SavingReadResponse{
		Date:       tranferRead.Date,
		SavingType: tranferRead.SavingType,
		Flow:       tranferRead.Flow,
		Amount:     tranferRead.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferReadResponse)
}

func (h *HttpSavingHandler) UpdateSaving(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := SavingRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tranfer := domain.Saving{
		Date:         body.Date,
		SavingTypeID: body.SavingTypeID,
		Flow:         body.Flow,
		Amount:       body.Amount,
	}

	var updatedSaving domain.Saving
	updatedSaving, err = h.tranferUsecase.Update(id, tranfer)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := SavingResponse{
		ID:           updatedSaving.ID,
		Date:         updatedSaving.Date,
		SavingTypeID: updatedSaving.SavingTypeID,
		Flow:         updatedSaving.Flow,
		Amount:       updatedSaving.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferResponse)
}

func (h *HttpSavingHandler) DeleteSaving(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var deletedSaving domain.Saving
	deletedSaving, err = h.tranferUsecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := SavingResponse{
		ID:           deletedSaving.ID,
		Date:         deletedSaving.Date,
		SavingTypeID: deletedSaving.SavingTypeID,
		Flow:         deletedSaving.Flow,
		Amount:       deletedSaving.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferResponse)
}
