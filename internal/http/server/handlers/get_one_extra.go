package handlers

import (
	"holidays-seeker/internal/core/usecases"
	"holidays-seeker/internal/infraestucture/dependencies"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FindOneHolidayHandler struct {
	uc *usecases.FinderHolidaysByExtra
}

func NewFindOneHolidayHandler(container dependencies.Container) *FindOneHolidayHandler {
	return &FindOneHolidayHandler{
		uc: usecases.NewFinderHolidaysByExtra(container.HolidaysRepository()),
	}
}

func (handler *FindOneHolidayHandler) FindOneHoliday(ctx *fiber.Ctx) error {
	// extraName := ctx.Params("extra")
	extraName := ctx.Query("extra")

	// TODO: Validate extraName
	holiday, err := handler.uc.Execute(ctx.Context(), extraName)
	if err != nil {
		ctx.JSON(err.Error())
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}
	ctx.JSON(holiday)
	ctx.SendStatus(http.StatusOK)
	return nil
}
