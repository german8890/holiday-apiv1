package inmemory

import (
	"context"
	"errors"
	"holidays-seeker/cmd/config"
	"holidays-seeker/internal/core/domain/fetcher"
	"holidays-seeker/internal/core/domain/holiday"
	"holidays-seeker/internal/http/client/fetcherService"
)

type Repository struct {
	list    []holiday.Holiday
	fetcher fetcher.Service
}

func (r *Repository) Find(ctx context.Context, date string) ([]holiday.Holiday, error) {
	if date != "" {
		var locals []holiday.Holiday
		for _, v := range r.list {
			if v.Date == date {
				locals = append(locals, v)
			}
		}
		return locals, nil
	}

	return r.list, nil
}

func (r *Repository) FindOne(ctx context.Context, extra string) (holiday.Holiday, error) {
	if extra == "" {
		return holiday.Holiday{}, errors.New("extra is required")
	}
	for _, v := range r.list {
		if v.Extra == extra {
			return v, nil
		}
	}
	return holiday.Holiday{}, nil
}

func (r *Repository) LoadAll(ctx context.Context, holidays []holiday.Holiday) error {
	r.list = holidays
	return nil
}

func NewInMemoryRepository(c config.Config) holiday.Repository {
	return &Repository{
		list:    []holiday.Holiday{},
		fetcher: fetcherService.NewFetcherService(c),
	}
}
