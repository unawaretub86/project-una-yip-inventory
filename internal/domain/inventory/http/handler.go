package http

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/usecase"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

const (
	suffixErr  = "Error"
	suffixInv = "Inventory"
)


type (
	Handler struct {
		UseCase usecase.UseCase
	}
)

func NewHandler(container *dependencies.Container) *Handler {
	return &Handler{
		UseCase: usecase.NewUse(container),
	}
}
