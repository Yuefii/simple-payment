package handlers

import (
	"net/http"
	"simple-payment/config"
	"simple-payment/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.GetDB()
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	db := config.GetDB()
	db.Find(&products)

	c.JSON(http.StatusOK, products)
}

func GetProductsByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	db := config.GetDB()
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	db := config.GetDB()
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	db := config.GetDB()
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
