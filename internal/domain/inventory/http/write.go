package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/project-una-yip-inventory/internal/domain/inventory/entities"
	"github.com/unawaretub86/project-una-yip-inventory/utils"
)

func (handler Handler) CreateItem(c *gin.Context) {
	itemData := &entities.TechItem{}

	if err := c.ShouldBindJSON(&itemData); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	item, err := handler.UseCase.CreateItem(itemData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusCreated, suffixInv, item)
}

func (handler Handler) UpdateItem(c *gin.Context) {
	itemData := &entities.TechItem{}

	if err := c.ShouldBindUri(&itemData); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	if err := c.ShouldBindJSON(&itemData); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	item, err := handler.UseCase.UpdateItem(itemData.ID, itemData)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusCreated, suffixInv, item)
}
