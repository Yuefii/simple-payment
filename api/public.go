package api

import (
	"simple-payment/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func PublicRouter(r *gin.Engine) {
	r.POST("auth/register", handlers.Register)
	r.POST("auth/login", handlers.Login)

	r.GET("/api/products", handlers.GetProducts)
	r.GET("/api/products/:id", handlers.GetProductsByID)
}
