package handlers

import (
	"encoding/xml"
	"holidays-seeker/internal/core/domain/holiday"
	"holidays-seeker/internal/core/usecases"
	"holidays-seeker/internal/infraestucture/dependencies"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FindDateHolidaysHandler struct {
	uc *usecases.FinderDateHolidays // .FinderAllHolidays
}

func NewFinDateHolidaysHandler(container dependencies.Container) *FindDateHolidaysHandler {
	return &FindDateHolidaysHandler{
		uc: usecases.NewFinderDateHolidays(container.HolidaysRepository()),
	}
}

// swagger: route GET /holidays getAllHolidays
//
// # Get all holidays
//
// Responses:
//
// - 200: GetDateHolidaysResponse
func (handler *FindDateHolidaysHandler) GetDateHolidays(ctx *fiber.Ctx) error {
	date := ctx.Query("date", "")
	responseType := ctx.Query("type")

	holidays, err := handler.uc.Execute(ctx.Context(), date)

	if responseType == "xml" {
		var holidaysXML holiday.Holidays

		holidaysXML.Holidays = holidays

		xmlBytes, err := xml.Marshal(holidaysXML)
		if err != nil {
			return err
		}
		ctx.Set("Content-Type", "application/xml")
		return ctx.SendString(string(xmlBytes))
	}

	if err != nil {
		return reply(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	return reply(ctx, http.StatusOK, "OK", holidays)
}
