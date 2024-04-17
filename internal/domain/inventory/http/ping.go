package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/project-una-yip-inventory/utils"
)

func (handler Handler) Ping(c *gin.Context) {
	utils.EndWithStatus(c, http.StatusOK, suffixInv, "pong")
}