package usecase

import "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"

func (useCase useCase) GetInventory() (*entities.Inventory, error) {
	return useCase.repo.GetInventory()
}