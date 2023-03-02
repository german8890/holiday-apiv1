package dependencies

import (
	"holidays-seeker/cmd/config"
	"holidays-seeker/internal/core/domain/fetcher"
	"holidays-seeker/internal/core/domain/holiday"
	"holidays-seeker/internal/http/client/fetcherService"
	"holidays-seeker/internal/infraestucture/storage"
)

type Container interface {
	Config() config.Config
	FetcherService() fetcher.Service
	HolidaysRepository() holiday.Repository
}

type container struct {
	cfg                config.Config
	fetcherService     fetcher.Service
	holidaysRepository holiday.Repository
}

func NewContainer(c config.Config) Container {
	fSvc := fetcherService.NewFetcherService(c)

	return &container{
		cfg:                c,
		holidaysRepository: storage.New(c),
		fetcherService:     fSvc,
	}
}

func (c *container) Config() config.Config {
	return c.cfg
}

func (c *container) FetcherService() fetcher.Service {
	return c.fetcherService
}

func (c *container) HolidaysRepository() holiday.Repository {
	return c.holidaysRepository
}
