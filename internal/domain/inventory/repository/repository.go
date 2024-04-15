package repository

import (
	database "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/repository/databse"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type (
	Repo interface {
	}

	repository struct {
		database database.Database
	}
)

func NewRepository(container *dependencies.Container) Repo {
	return &repository{
		database: database.NewDatabase(container),
	}
}