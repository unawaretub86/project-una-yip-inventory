package usecase

import "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"

func (useCase useCase) GetInventory() (*entities.Inventory, error) {
	return useCase.repo.GetInventory()
}

func (useCase useCase) CreateItem(item *entities.TechItem) (*entities.TechItem, error) {
	return useCase.repo.CreateItem(item)
}

func (useCase useCase) UpdateItem(id int64, item *entities.TechItem) (*entities.TechItem, error) {
	return useCase.repo.UpdateItem(id, item)
}

func (useCase useCase) GetItemById(id int64) (*entities.TechItem, error)  {
	return useCase.repo.GetItemById(id)
}
