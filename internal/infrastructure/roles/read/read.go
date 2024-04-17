package read

import (
	"github.com/gin-gonic/gin"

	inventoryHandler "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/http"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type read struct {
	container *dependencies.Container
}

func NewRead(container *dependencies.Container) *read {
	return &read{
		container: container,
	}
}

func (read *read) RegisterRoutes(basePath string, r *gin.Engine) {
	invHandler := inventoryHandler.NewHandler(read.container)

	r.GET("/ping", invHandler.Ping)

	v1Group := r.Group(basePath + "/v1")

	v1Group.GET("/inventory", invHandler.GetInventory)
}
