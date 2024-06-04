package main

import (
	"simple-payment/api"
	"simple-payment/config"
	"simple-payment/pkg/models"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	r.Use(cors.New(config))

	api.PublicRouter(r)
	api.PrivateRouter(r)

	r.Run(":8000")
}
