package server

import (
	"context"
	"fmt"
	"holidays-seeker/internal/http/server/handlers"
	"holidays-seeker/internal/infraestucture/dependencies"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ServerHTTP struct{}

func Run(container dependencies.Container) {
	r := fiber.New()

	// fill data
	list, err := container.FetcherService().RetrieveHolidays(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	container.HolidaysRepository().LoadAll(context.TODO(), list)
	api := r.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	getAllHandler := handlers.NewFinDateHolidaysHandler(container)
	getOneHandler := handlers.NewFindOneHolidayHandler(container)

	v1.Get("/holidays/date", getAllHandler.GetDateHolidays)
	v1.Get("/holidays/extra", getOneHandler.FindOneHoliday)

	port := container.Config().Server.Port

	log.Fatal(r.Listen(fmt.Sprintf(":%s", port)))
}
