package database_test

import (
	"testing"

	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/repository/database"
	mock_database "github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/repository/database/mocks"
	"go.uber.org/mock/gomock"
)

func GetInventorySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	
    defer ctrl.Finish()

	mockDB := mock_database.NewMockDatabase(ctrl)

	db := &data{
		dataBase: mockDB,
	}

}