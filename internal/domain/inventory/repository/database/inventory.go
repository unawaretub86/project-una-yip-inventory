package database

import (
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"
	"gorm.io/gorm"
)

func (database DatabaseInventory) GetInventory() (*entities.Inventory, error) {
	Inventory := &entities.Inventory{}

	result := database.dataBase.Connection().Find(&Inventory)

	if result.Error != nil {
		return nil, result.Error
	}

	return Inventory, nil
}

func (database DatabaseInventory) GetItemById(id int64) (*entities.TechItem, error)  {
	result, err := database.getItem(id)
	if err.Error != nil {
		return nil, err.Error
	}

	return result, nil
}

func (database DatabaseInventory) CreateItem(item *entities.TechItem) (*entities.TechItem, error) {
	result := database.dataBase.Connection().Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (database DatabaseInventory) UpdateItem(id int64, item *entities.TechItem) (*entities.TechItem, error) {
	resultData, err := database.getItem(id)
	if err.Error != nil {
		return nil, err.Error
	}

	item.ID = resultData.ID

	result := database.dataBase.Connection().
	Where("id = ?", item.ID).
	Updates(item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (database DatabaseInventory) getItem(id int64) (*entities.TechItem, *gorm.DB)  {
	item := &entities.TechItem{}

	result := database.dataBase.Connection().
	Where("id = ?", id).
	Take(&item)

	return item, result
}
