package usecases

import (
	"context"
	"holidays-seeker/internal/core/domain/holiday"
)

// FinderAllHolidays is the use case than find all holiday
type FinderDateHolidays struct {
	holidaysRepository holiday.Repository
}

func NewFinderDateHolidays(repository holiday.Repository) *FinderDateHolidays {
	return &FinderDateHolidays{
		repository,
	}
}

// Execute finder in the repository of holiday
func (f *FinderDateHolidays) Execute(ctx context.Context, date string) ([]holiday.Holiday, error) {
	if err := holiday.ValidateHolidayDate(date); err != nil {
		return nil, err
	}
	holidaysList, err := f.holidaysRepository.Find(ctx, date)
	if err != nil {
		return nil, err
	}
	return holidaysList, nil
}
