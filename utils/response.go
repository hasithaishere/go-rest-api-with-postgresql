package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context, status int, payload interface{}) {
    c.JSON(status, payload)
}
