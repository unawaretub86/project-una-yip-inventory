package database

import (
	database "github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/configuration/database"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type (
	Database interface {
	}

	databaseDomain struct {
		db database.Database
	}
)

func NewDatabase(container *dependencies.Container) Database {
	return &databaseDomain{
		db: container.Database(),
	}
}