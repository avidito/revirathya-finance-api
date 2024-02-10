package expense

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoute(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/expenses")
	routes.Post("/", h.AddExpense)
	routes.Get("/", h.GetExpenses)
	routes.Get("/:id", h.GetExpense)
	routes.Put("/:id", h.UpdateExpense)
	routes.Delete("/:id", h.DeleteExpense)
}
