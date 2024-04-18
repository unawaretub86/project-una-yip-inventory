package write

import (
	"github.com/gin-gonic/gin"
	
	inventoryHandler "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/http"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

type write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *write {
	return &write{
		container: container,
	}
}

func (write *write) RegisterRoutes(basePath string, r *gin.Engine) {
	invHandler := inventoryHandler.NewHandler(write.container)
	
	r1 := r.Group(basePath + "/v1")

	r1.POST("/item", invHandler.CreateItem)
	r1.PATCH("/item/:id", invHandler.UpdateItem)
}
