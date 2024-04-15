package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	DBMaxIdleConns      = 25
	DBMaxOpenConns      = 50
	ConnMaxLifetimeSecs = 300
	DBReadTimeout       = 10
	DBWriteTimeout      = 10
	DBTimeout           = 10
)

type (
	Database interface {
		Connection() *gorm.DB
	}

	service struct {
		db *gorm.DB
	}
)

func ConnectDB(connectionString string) (Database, error) {

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to db")
	}

	err = db.Use(GetPlugin())
	if err != nil {
		return nil, err
	}

	service := &service{
		db: db,
	}

	return service, nil
}

func (service *service) Connection() *gorm.DB {
	return service.db
}

func GetPlugin() gorm.Plugin {
	connMaxLifetime := ConnMaxLifetimeSecs * time.Second

	return dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(connMaxLifetime).
		SetMaxIdleConns(DBMaxIdleConns).
		SetMaxOpenConns(DBMaxOpenConns)
}
