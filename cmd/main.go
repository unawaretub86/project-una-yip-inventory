package main

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/configuration/server"
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/dependencies"
)

const basePath = ""

func main() {
	container := dependencies.StartDependencies()

	server.NewServer(basePath, container)
}