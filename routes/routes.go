package routes

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mryan-3/smart/db"
	"github.com/mryan-3/smart/models"
)

var validate = validator.New()

func CreateSmartData(c *gin.Context) {
	var smart_data models.Smart_Data

	err := c.ShouldBindJSON(&smart_data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

    if err := validate.Struct(smart_data); err != nil {
        errs := err.(validator.ValidationErrors)
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": errs,
        })
    }

	var data []models.Smart_Data

	dbErr := db.Supabase.DB.From("Data").Insert(&smart_data).Execute(&data)

	if dbErr != nil {
		c.AbortWithError(http.StatusInternalServerError, dbErr)
		return
	}

	c.JSON(http.StatusCreated, data[0])
}

func GetData(c *gin.Context) {
	var data []models.Smart_Data

	err := db.Supabase.DB.From("Data").Select("*").Execute(&data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func getDataById(c *gin.Context) {
	var data []models.Smart_Data

	id := c.Param("meter-id")

	err := db.Supabase.DB.From("Data").Select("*").Eq("meter-id", id).Execute(&data)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
