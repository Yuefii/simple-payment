package api

import (
	"simple-payment/pkg/handlers"
	"simple-payment/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PrivateRouter(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/api/products", handlers.CreateProduct)
		authorized.PATCH("/api/products/:id", handlers.UpdateProduct)
		authorized.DELETE("/api/products/:id", handlers.DeleteProduct)

		authorized.POST("api/orders", handlers.CreateOrder)
		authorized.GET("api/orders/:id/qrcode", handlers.GetOrderByQRCode)
	}
}
