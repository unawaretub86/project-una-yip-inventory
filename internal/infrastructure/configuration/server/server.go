package server

import (
	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/roles/read"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/roles/write"
)

func NewServer(basePath string, container *dependencies.Container) {
	r := gin.Default()

	read.NewRead(container).RegisterRoutes(basePath, r)
	write.NewWrite(container).RegisterRoutes(basePath, r)

	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}