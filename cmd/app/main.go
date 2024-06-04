package main

import (
	"simple-payment/api"
	"simple-payment/config"
	"simple-payment/pkg/models"

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

	api.PublicRouter(r)
	api.PrivateRouter(r)

	r.Run(":8000")
}
