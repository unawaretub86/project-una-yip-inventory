package repository

import "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"

func (repository repository) GetInventory() (*entities.Inventory, error) {
	return repository.database.GetInventory()
}