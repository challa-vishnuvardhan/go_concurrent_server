package handler

import (
	"encoding/json"
	"go_concurrent_server/resource"
	"go_concurrent_server/service"

	"github.com/gin-gonic/gin"
)

func ServiceClient(c *gin.Context) {

	var input resource.Request

	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Payload"})
		return
	}

	result, err := service.ServiceClient(input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}
