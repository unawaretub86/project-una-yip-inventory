package read

import (
	"github.com/gin-gonic/gin"

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
}
