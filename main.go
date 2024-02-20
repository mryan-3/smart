package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Smart_Data struct {
	Timestamp   string `json:"timestamp"`
	Meter_id    string `json:"meter_id"`
	Consumption string `json:"consumption"`
}

var whole_data = []Smart_Data{}

func main() {
	fmt.Println("hello world jones")
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, whole_data)
	})

	r.GET("/api/v1/smart-meter-data/:meter_id", func(c *gin.Context) {
		smart_meter_id := c.Param("meter_id")
		var data Smart_Data

		for _, d := range whole_data {
			if d.Meter_id == smart_meter_id {
				data = d
			}
		}
		if data.Meter_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Meter ID not found",
			})
		}

		c.JSON(http.StatusOK, data)
	})

	r.DELETE("/api/v1/smart_meter_data/meter_id", func(c *gin.Context) {
		smart_meter_id := c.Param("meter_id")

		for i, d := range whole_data {
			if d.Meter_id == smart_meter_id {
				whole_data = append(whole_data[:i], whole_data[i+1:]...)
				break
			}
		}
		c.JSON(200, gin.H{"message": "Meter Deleted"})
	})

	r.POST("/books", func(c *gin.Context) {
		var smart_data Smart_Data

		if err := c.ShouldBindJSON(&smart_data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		whole_data = append(whole_data, smart_data)

		c.JSON(http.StatusCreated, smart_data)
	})
	r.Run()
}
