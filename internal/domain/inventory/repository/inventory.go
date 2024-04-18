package repository

import "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"

func (repository repository) GetInventory() (*entities.Inventory, error) {
	return repository.database.GetInventory()
}

func (repository repository) CreateItem(item *entities.TechItem) (*entities.TechItem, error) {
	return repository.database.CreateItem(item)
}

func (repository repository) UpdateItem(id int64,item *entities.TechItem) (*entities.TechItem, error) {
	return repository.database.UpdateItem(id, item)
}