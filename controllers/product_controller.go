package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindProducts(c *gin.Context) {
    var products []models.Product
    models.DB.Find(&products)
    utils.RespondJSON(c, http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product := models.Product{Name: input.Name, Price: input.Price}
    models.DB.Create(&product)
    utils.RespondJSON(c, http.StatusCreated, product)
}

func FindProduct(c *gin.Context) {
    var product models.Product
    if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Product not found!"})
        return
    }
    utils.RespondJSON(c, http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Product not found!"})
        return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&product).Updates(input)
    utils.RespondJSON(c, http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    var product models.Product
    if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Product not found!"})
        return
    }
    models.DB.Delete(&product)
    utils.RespondJSON(c, http.StatusNoContent, nil)
}
