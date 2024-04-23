package database

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"
	db "github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/configuration/database"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type (
	Database interface {
		GetInventory() (*entities.Inventory, error)
		GetItemById(int64) (*entities.TechItem, error)

		CreateItem(*entities.TechItem) (*entities.TechItem, error)
		UpdateItem(int64, *entities.TechItem) (*entities.TechItem, error)
	}

	DatabaseInventory struct {
		dataBase db.Database
	}
)

func NewDatabase(container *dependencies.Container) Database {
	return &DatabaseInventory{
		dataBase: container.Database(),
	}
}
