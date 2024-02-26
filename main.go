package main

import (
	//"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/gin-contrib/cors"
	"github.com/mryan-3/smart/db"
	"github.com/mryan-3/smart/routes"
    "log"
    "os"
)


func main() {
	fmt.Println("hello world jones")
	db.CreateClient()
	r := gin.New()
    r.Use(cors.Default())



	r.POST("/api/v1/smart-meter-data", routes.CreateSmartData)
	r.GET("/api/v1/smart-meter-data", routes.GetData)
	r.GET("/api/v1/smart-meter-data/:meter-id", routes.GetDataById)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    if err := r.Run(":" + port); err != nil {
        log.Panicf("error: %s", err)
    }
	r.Run()
}
