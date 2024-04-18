package database

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"
	"gorm.io/gorm"
)

func (database database) GetInventory() (*entities.Inventory, error) {
	Inventory := &entities.Inventory{}

	result := database.db.Connection().Find(&Inventory)

	if result.Error != nil {
		return nil, result.Error
	}

	return Inventory, nil
}

func (database database) CreateItem(item *entities.TechItem) (*entities.TechItem, error) {
	result := database.db.Connection().Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (database database) UpdateItem(id int64, item *entities.TechItem) (*entities.TechItem, error) {
	resultData, err := database.getItem(id)
	if err.Error != nil {
		return nil, err.Error
	}

	item.ID = resultData.ID

	result := database.db.Connection().
	Where("id = ?", item.ID).
	Updates(item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (database database) getItem(id int64) (*entities.TechItem, *gorm.DB)  {
	item := &entities.TechItem{}

	result := database.db.Connection().
	Where("id = ?", id).
	Take(&item)

	return item, result
}
