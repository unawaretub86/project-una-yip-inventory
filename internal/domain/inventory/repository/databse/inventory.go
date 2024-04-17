package database

import "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"

func (database database) GetInventory() (*entities.Inventory, error) {
	Inventory := &entities.Inventory{}

	result := database.db.Connection().Find(&Inventory)

	if result.Error != nil {
		return nil, result.Error
	}

	return Inventory, nil
}