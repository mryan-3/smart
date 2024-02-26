package main

import (
	//"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
    "github.com/mryan-3/smart/routes"
    "github.com/mryan-3/smart/db"
)



// Function to validate API key
//func validateAPIKey(c *gin.Context) (bool, string, error) {
	//apiKey := c.GetHeader("Authorization")
	//if apiKey == "" {
		//return false, "", errors.New("Missing Key")
	//}
	//role, ok := db[apiKey]

	//if !ok {
		//return false, "", errors.New("Invalid Key")
	//}
	//return true, role, nil
//}

//func apiKeyAuthMiddleware() gin.HandlerFunc {
	//return func(c *gin.Context) {
		//isValid, role, err := validateAPIKey(c)
		//if err != nil {
			//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				//"error": err.Error(),
			//})
		//}

		//if !isValid {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				//"error": "Invalid API Key",
			//})
			//return
		//}

		//c.Set("UserRole", role)
		//c.Next()
	//}

//}


func main() {
	fmt.Println("hello world jones")
    db.CreateClient()
	r := gin.New()
    //r.Use(apiKeyAuthMiddleware())
    r.POST("/api/v1/smart-meter-data", routes.CreateSmartData)
    r.GET("/api/v1/smart-meter-data", routes.GetData)
    r.GET("/api/v1/smart-meter-data/:meter-id", routes.GetDataById)
	r.Run()
}
