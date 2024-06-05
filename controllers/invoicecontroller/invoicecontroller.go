package invoicecontroller

import (
	"net/http"

	"github.com/marwenbhriz/go-payment-api/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var invoices []models.Invoice

	models.DB.Find(&invoices)
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

func Show(c *gin.Context) {
	var invoice models.Invoice
	id := c.Param("id")

	if err := models.DB.First(&invoice, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"invoice": invoice})
}
