package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Smart_Data struct {
	Timestamp   string `json:"timestamp"`
	Meter_id    string `json:"meter_id"`
	Consumption string `json:"consumption"`
}

var whole_data = []Smart_Data{}
var db map[string]string

// Function to validate API key
func validateAPIKey(c *gin.Context) (bool, string, error) {
	apiKey := c.GetHeader("Authorization")
	if apiKey == "" {
		return false, "", errors.New("Missing Key")
	}
	role, ok := db[apiKey]

	if !ok {
		return false, "", errors.New("Invalid Key")
	}
	return true, role, nil
}

func apiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValid, role, err := validateAPIKey(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API Key",
			})
			return
		}

		c.Set("UserRole", role)
		c.Next()
	}

}


func main() {
	fmt.Println("hello world jones")
	r := gin.New()
    r.Use(apiKeyAuthMiddleware())
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

	r.Run()
}
