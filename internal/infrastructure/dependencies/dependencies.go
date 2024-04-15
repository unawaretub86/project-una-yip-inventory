package dependencies

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/configuration"

	"github.com/unawaretub86/project-una-yip-inventory/internal/infrastructure/configuration/database"
)

type Container struct {
	database database.Database
}

func StartDependencies() *Container {
	dbUrl := configuration.GetDatabaseConfig()

	db, err := database.ConnectDB(dbUrl)
	if err != nil {
		panic(err.Error())
	}

	container := &Container{
		database: db,
	}

	return container
}

func (container *Container) Database() database.Database {
	return container.database
}
