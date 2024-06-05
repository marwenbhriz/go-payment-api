package usercontroller

import (
	"net/http"

	"github.com/marwenbhriz/go-payment-api/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
