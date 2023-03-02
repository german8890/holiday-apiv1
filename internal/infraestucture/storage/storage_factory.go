package storage

import (
	"holidays-seeker/cmd/config"
	"holidays-seeker/internal/core/domain/holiday"
	"holidays-seeker/internal/infraestucture/storage/inmemory"
)

func New(c config.Config) holiday.Repository {
	return inmemory.NewInMemoryRepository(c)
}
