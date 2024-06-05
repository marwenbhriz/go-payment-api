package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marwenbhriz/go-payment-api/controllers/apicontroller"
	"github.com/marwenbhriz/go-payment-api/controllers/companycontroller"
	"github.com/marwenbhriz/go-payment-api/controllers/invoicecontroller"
	"github.com/marwenbhriz/go-payment-api/controllers/usercontroller"
	"github.com/marwenbhriz/go-payment-api/models"
)

func main() {

	router := gin.Default()
	// log.Println("Payment API start listen on port 8080.")
	models.ConnectDatabase()

	router.GET("/api", apicontroller.Index)

	// companies routes
	router.GET("/api/companies", companycontroller.Index)
	router.GET("/api/company/:id", companycontroller.Show)
	router.POST("/api/company", companycontroller.Create)
	router.PUT("/api/company/:id", companycontroller.Update)
	router.DELETE("/api/company", companycontroller.Delete)

	// users routes
	router.GET("/api/users", usercontroller.Index)
	router.GET("/api/users/:id", usercontroller.Show)

	// invoices routes
	router.GET("/api/invoices", invoicecontroller.Index)
	router.GET("/api/invoices/:id", invoicecontroller.Show)
	router.POST("/api/invoices", apicontroller.Index)

	router.Run(":8080")
}
