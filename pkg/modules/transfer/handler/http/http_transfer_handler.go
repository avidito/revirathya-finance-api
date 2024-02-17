package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpTransferHandler struct {
	tranferUsecase domain.TransferUsecase
}

func NewHttpTransferHandler(app *fiber.App, u domain.TransferUsecase) {
	h := &HttpTransferHandler{
		tranferUsecase: u,
	}

	routes := app.Group("/transfers")
	routes.Post("/", h.CreateTransfer)
	routes.Get("/", h.FetchTransfers)
	routes.Get("/:id", h.GetTransferByID)
	routes.Put("/:id", h.UpdateTransfer)
	routes.Delete("/:id", h.DeleteTransfer)
}

// Handler
func (h *HttpTransferHandler) CreateTransfer(c *fiber.Ctx) error {
	var err error

	body := TransferRequestBody{}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tranfer := domain.Transfer{
		Date:          body.Date,
		SourceID:      body.SourceID,
		DestinationID: body.DestinationID,
		Description:   body.Description,
		Amount:        body.Amount,
	}

	var createdTransfer domain.Transfer
	createdTransfer, err = h.tranferUsecase.Create(tranfer)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := TransferResponse{
		ID:            createdTransfer.ID,
		Date:          createdTransfer.Date,
		SourceID:      createdTransfer.SourceID,
		DestinationID: createdTransfer.DestinationID,
		Description:   createdTransfer.Description,
		Amount:        createdTransfer.Amount,
	}
	return c.Status(fiber.StatusCreated).JSON(&tranferResponse)
}

func (h *HttpTransferHandler) FetchTransfers(c *fiber.Ctx) error {
	var date string
	var source string
	var destination string
	var err error

	q := c.Queries()
	date = q["date"]
	source = q["source"]
	destination = q["destination"]

	var tranferReadList []domain.TransferRead
	tranferReadList, err = h.tranferUsecase.Fetch(date, source, destination)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var tranferReadListResponse []TransferReadResponse
	for _, tranferRead := range tranferReadList {
		tranferReadListResponse = append(
			tranferReadListResponse,
			TransferReadResponse{
				Date:        tranferRead.Date,
				Source:      tranferRead.Source,
				Destination: tranferRead.Destination,
				Description: tranferRead.Description,
				Amount:      tranferRead.Amount,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&tranferReadListResponse)
}

func (h *HttpTransferHandler) GetTransferByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var tranferRead domain.TransferRead
	tranferRead, err = h.tranferUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferReadResponse := TransferReadResponse{
		Date:        tranferRead.Date,
		Source:      tranferRead.Source,
		Destination: tranferRead.Destination,
		Description: tranferRead.Description,
		Amount:      tranferRead.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferReadResponse)
}

func (h *HttpTransferHandler) UpdateTransfer(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := TransferRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tranfer := domain.Transfer{
		Date:          body.Date,
		SourceID:      body.SourceID,
		DestinationID: body.DestinationID,
		Description:   body.Description,
		Amount:        body.Amount,
	}

	var updatedTransfer domain.Transfer
	updatedTransfer, err = h.tranferUsecase.Update(id, tranfer)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := TransferResponse{
		ID:            updatedTransfer.ID,
		Date:          updatedTransfer.Date,
		SourceID:      updatedTransfer.SourceID,
		DestinationID: updatedTransfer.DestinationID,
		Description:   updatedTransfer.Description,
		Amount:        updatedTransfer.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferResponse)
}

func (h *HttpTransferHandler) DeleteTransfer(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var deletedTransfer domain.Transfer
	deletedTransfer, err = h.tranferUsecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	tranferResponse := TransferResponse{
		ID:            deletedTransfer.ID,
		Date:          deletedTransfer.Date,
		SourceID:      deletedTransfer.SourceID,
		DestinationID: deletedTransfer.DestinationID,
		Description:   deletedTransfer.Description,
		Amount:        deletedTransfer.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&tranferResponse)
}
