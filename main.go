package main

import (
	"simple-payment/config"
	"simple-payment/handlers"
	"simple-payment/middleware"
	"simple-payment/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Connect()
	db := config.GetDB()
	defer db.Close()

	models.Migrate(db)

	r := gin.Default()

	r.POST("auth/register", handlers.Register)
	r.POST("auth/login", handlers.Login)

	r.GET("/api/products", handlers.GetProducts)
	r.GET("/api/products/:id", handlers.GetProductsByID)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/api/products", handlers.CreateProduct)
		authorized.PATCH("/api/products/:id", handlers.UpdateProduct)
		authorized.DELETE("/api/products/:id", handlers.DeleteProduct)

		authorized.POST("api/orders", handlers.CreateOrder)
		authorized.GET("api/orders/:id/qrcode", handlers.GetOrderByQRCode)
	}

	r.Run()
}
