package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/project-una-yip-inventory/utils"
)

func (handler Handler) GetInventory(c *gin.Context) {
	inventory, err := handler.UseCase.GetInventory()
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixInv, inventory)
}

func (handler Handler) Ping(c *gin.Context) {
	utils.EndWithStatus(c, http.StatusOK, suffixInv, "pong")
}
