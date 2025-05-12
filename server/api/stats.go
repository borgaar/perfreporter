package api

import (
	sensors "github.com/borgaar/perfreporter/lib"
	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	data, err := sensors.ReadSensorData()

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, data)
}
