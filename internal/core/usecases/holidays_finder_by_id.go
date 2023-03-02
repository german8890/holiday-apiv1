package usecases

import (
	"context"
	"errors"
	"holidays-seeker/internal/core/domain/holiday"
)

// FinderHolidaysByExtra is the use case than find all holidays
type FinderHolidaysByExtra struct {
	holidaysRepository holiday.Repository
}

func NewFinderHolidaysByExtra(repository holiday.Repository) *FinderHolidaysByExtra {
	return &FinderHolidaysByExtra{
		repository,
	}
}

// Execute finder a holiday by his ID in the repository of holidays
func (pf *FinderHolidaysByExtra) Execute(ctx context.Context, extra string) (holiday.Holiday, error) {
	if extra == "" {
		return holiday.Holiday{}, errors.New("extra is required")
	}
	if err := holiday.ValidateHolidayExtra(extra); err != nil {
		return holiday.Holiday{}, err
	}
	result, err := pf.holidaysRepository.FindOne(ctx, extra)
	if err != nil {
		return holiday.Holiday{}, err
	}
	return result, nil
}
