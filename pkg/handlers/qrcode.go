package handlers

import (
	"fmt"
	"net/http"
	"simple-payment/config"
	"simple-payment/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func GetOrderByQRCode(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order

	db := config.GetDB()
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	qrContent := fmt.Sprintf("Order ID: %d, User ID: %d, Product ID: %d, Quantity: %d, Total: %.2f", order.ID, order.UserID, order.ProductID, order.Quantity, order.Total)
	qr, err := qrcode.New(qrContent, qrcode.Medium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(http.StatusOK)
	err = qr.Write(256, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write QR code"})
		return
	}
}
