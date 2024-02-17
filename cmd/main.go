package main

import (
	"log"

	"github.com/avidito/revirathya-finance-api/pkg"
	"github.com/avidito/revirathya-finance-api/pkg/common/config"
	"github.com/avidito/revirathya-finance-api/pkg/common/db"
	"github.com/avidito/revirathya-finance-api/seeds"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	// Seeder
	seedData := seeds.NewSeedData()
	seeder := seeds.NewSeeder(db, seedData)
	seeder.Load()

	// Base Endpoint
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	// Layering
	repository := pkg.NewRepository(db)
	usecase := pkg.NewUsecase(repository)
	pkg.NewHandler(app, usecase)

	app.Listen(c.Port)
}
