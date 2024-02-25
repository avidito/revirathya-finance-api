package http

import (
	"strconv"

	"github.com/avidito/revirathya-finance-api/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Define
type HttpReferenceHandler struct {
	referenceUsecase domain.ReferenceUsecase
}

func NewHttpReferenceHandler(app *fiber.App, u domain.ReferenceUsecase) {
	h := &HttpReferenceHandler{
		referenceUsecase: u,
	}

	routes := app.Group("/reference")
	routes.Get("/income_types", h.FetchIncomeTypes)
	routes.Get("/locations", h.FetchLocations)
	routes.Get("/budget_groups", h.FetchBudgetGroups)
	routes.Get("/budget_types", h.FetchBudgetTypes)
	routes.Get("/saving_types", h.FetchSavingTypes)
	routes.Get("/income_types/:id", h.GetIncomeTypeByID)
	routes.Get("/locations/:id", h.GetLocationByID)
	routes.Get("/budget_groups/:id", h.GetBudgetGroupByID)
	routes.Get("/budget_types/:id", h.GetBudgetTypeByID)
	routes.Get("/saving_types/:id", h.GetSavingTypeByID)
}

// Handler
func (h *HttpReferenceHandler) GetIncomeTypeByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var refIncomeType domain.RefIncomeType
	refIncomeType, err = h.referenceUsecase.GetIncomeTypeByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	refIncomeTypeResponse := RefIncomeTypeResponse{
		ID:         refIncomeType.ID,
		IncomeType: refIncomeType.IncomeType,
	}
	return c.Status(fiber.StatusOK).JSON(&refIncomeTypeResponse)
}

func (h *HttpReferenceHandler) GetLocationByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var refLocation domain.RefLocation
	refLocation, err = h.referenceUsecase.GetLocationByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	refLocationResponse := RefLocationResponse{
		ID:       refLocation.ID,
		Location: refLocation.Location,
	}
	return c.Status(fiber.StatusOK).JSON(&refLocationResponse)
}

func (h *HttpReferenceHandler) GetBudgetGroupByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var refBudgetGroup domain.RefBudgetGroup
	refBudgetGroup, err = h.referenceUsecase.GetBudgetGroupByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	refBudgetGroupResponse := RefBudgetGroupResponse{
		ID:          refBudgetGroup.ID,
		BudgetGroup: refBudgetGroup.BudgetGroup,
	}
	return c.Status(fiber.StatusOK).JSON(&refBudgetGroupResponse)
}

func (h *HttpReferenceHandler) GetBudgetTypeByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var refBudgetTypeRead domain.RefBudgetTypeRead
	refBudgetTypeRead, err = h.referenceUsecase.GetBudgetTypeByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	refBudgetTypeReadResponse := RefBudgetTypeReadResponse{
		ID:          refBudgetTypeRead.ID,
		BudgetGroup: refBudgetTypeRead.BudgetGroup,
		BudgetType:  refBudgetTypeRead.BudgetType,
	}
	return c.Status(fiber.StatusOK).JSON(&refBudgetTypeReadResponse)
}

func (h *HttpReferenceHandler) GetSavingTypeByID(c *fiber.Ctx) error {
	var id int64
	var err error

	id, err = strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var refSavingType domain.RefSavingType
	refSavingType, err = h.referenceUsecase.GetSavingTypeByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	refSavingTypeResponse := RefSavingTypeResponse{
		ID:         refSavingType.ID,
		SavingType: refSavingType.SavingType,
	}
	return c.Status(fiber.StatusOK).JSON(&refSavingTypeResponse)
}

func (h *HttpReferenceHandler) FetchIncomeTypes(c *fiber.Ctx) error {
	var name string
	var err error

	q := c.Queries()
	name = q["name"]

	var refIncomeTypeList []domain.RefIncomeType
	refIncomeTypeList, err = h.referenceUsecase.FetchIncomeTypes(name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var refIncomeTypeListResponse []RefIncomeTypeResponse
	for _, refIncomeType := range refIncomeTypeList {
		refIncomeTypeListResponse = append(
			refIncomeTypeListResponse,
			RefIncomeTypeResponse{
				ID:         refIncomeType.ID,
				IncomeType: refIncomeType.IncomeType,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&refIncomeTypeListResponse)
}

func (h *HttpReferenceHandler) FetchLocations(c *fiber.Ctx) error {
	var name string
	var err error

	q := c.Queries()
	name = q["name"]

	var refLocationList []domain.RefLocation
	refLocationList, err = h.referenceUsecase.FetchLocations(name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var refLocationListResponse []RefLocationResponse
	for _, refLocation := range refLocationList {
		refLocationListResponse = append(
			refLocationListResponse,
			RefLocationResponse{
				ID:       refLocation.ID,
				Location: refLocation.Location,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&refLocationListResponse)
}

func (h *HttpReferenceHandler) FetchBudgetGroups(c *fiber.Ctx) error {
	var name string
	var err error

	q := c.Queries()
	name = q["name"]

	var refBudgetGroupList []domain.RefBudgetGroup
	refBudgetGroupList, err = h.referenceUsecase.FetchBudgetGroups(name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var refBudgetGroupListResponse []RefBudgetGroupResponse
	for _, refBudgetGroup := range refBudgetGroupList {
		refBudgetGroupListResponse = append(
			refBudgetGroupListResponse,
			RefBudgetGroupResponse{
				ID:          refBudgetGroup.ID,
				BudgetGroup: refBudgetGroup.BudgetGroup,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&refBudgetGroupListResponse)
}

func (h *HttpReferenceHandler) FetchBudgetTypes(c *fiber.Ctx) error {
	var name string
	var budget_group string
	var err error

	q := c.Queries()
	name = q["name"]
	budget_group = q["budget_group"]

	var refBudgetTypeReadList []domain.RefBudgetTypeRead
	refBudgetTypeReadList, err = h.referenceUsecase.FetchBudgetTypes(name, budget_group)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var refBudgetTypeReadListResponse []RefBudgetTypeReadResponse
	for _, refBudgetTypeRead := range refBudgetTypeReadList {
		refBudgetTypeReadListResponse = append(
			refBudgetTypeReadListResponse,
			RefBudgetTypeReadResponse{
				ID:          refBudgetTypeRead.ID,
				BudgetGroup: refBudgetTypeRead.BudgetGroup,
				BudgetType:  refBudgetTypeRead.BudgetType,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&refBudgetTypeReadListResponse)
}

func (h *HttpReferenceHandler) FetchSavingTypes(c *fiber.Ctx) error {
	var name string
	var err error

	q := c.Queries()
	name = q["name"]

	var refSavingTypeList []domain.RefSavingType
	refSavingTypeList, err = h.referenceUsecase.FetchSavingTypes(name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var refSavingTypeListResponse []RefSavingTypeResponse
	for _, refSavingType := range refSavingTypeList {
		refSavingTypeListResponse = append(
			refSavingTypeListResponse,
			RefSavingTypeResponse{
				ID:         refSavingType.ID,
				SavingType: refSavingType.SavingType,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&refSavingTypeListResponse)
}
