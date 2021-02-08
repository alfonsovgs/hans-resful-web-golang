package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(c *gin.Context) {
		//JSON serializer is availableon gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":8000")
}
