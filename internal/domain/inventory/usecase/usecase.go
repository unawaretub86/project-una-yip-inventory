package usecase

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/repository"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type (
	UseCase interface {
		GetInventory() (*entities.Inventory, error)
	}

	useCase struct {
		repo repository.Repo
	}
)

func NewUse(container *dependencies.Container) UseCase {
	return &useCase{
		repo: repository.NewRepository(container),
	}
}