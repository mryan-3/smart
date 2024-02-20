package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Smart_Data struct {
	Timestamp     string `json:"timestamp"`
	Meter_id  string `json:"meter_id"`
	Consumption string `json:"consumption"`
}

var whole_data = []Smart_Data{
}
func main() {
	fmt.Println("hello world jones")
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, whole_data)
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