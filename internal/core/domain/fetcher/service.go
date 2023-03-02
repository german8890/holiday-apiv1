package fetcher

import (
	"context"
	"holidays-seeker/internal/core/domain/holiday"
)

//go:generate mockgen -package mocks -destination mocks/holiday_service_mocks.go . Service

// Service is the service abstraction to Holiday
type Service interface {
	RetrieveHolidays(ctx context.Context) ([]holiday.Holiday, error)
}
