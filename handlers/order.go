package handlers

import (
	"fmt"
	"net/http"
	"simple-payment/config"
	"simple-payment/models"

	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	userIDFloat, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	userIDUint := uint(userIDFloat)

	var product models.Product
	db := config.GetDB()
	if err := db.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	total := float64(input.Quantity) * product.Price
	order := models.Order{
		UserID:    userIDUint,
		ProductID: input.ProductID,
		Quantity:  int(input.Quantity),
		Total:     total,
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Preload("User").Preload("Product").First(&order, order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qrCodeURL := fmt.Sprintf("/orders/%d/qrcode", order.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
		"qr_code": qrCodeURL,
	})
}
