package routes

import (
	"fmt"
	"go_concurrent_server/constants"
	"go_concurrent_server/handler"

	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()

	r.POST("/message", func(c *gin.Context) {
		go handler.ServiceClient(c)
	})

	err := r.Run(":" + constants.Port)
	if err != nil {
		fmt.Printf("Error listening on port %s: %v\n", constants.Port, err)
		panic(err)
	}

}
