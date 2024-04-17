package utils

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(err error, suffix string) gin.H {
	return gin.H{suffix: err.Error()}
}

func EndWithStatus(c *gin.Context, status int, suffix string, body any) {
	c.JSON(status, gin.H{suffix: body})
}

func EndWithStatusError(c *gin.Context, status int, suffix string, err error) {
	c.JSON(status, errorResponse(err, suffix))
}
