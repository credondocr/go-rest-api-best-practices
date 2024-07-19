package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/credondocr/go-rest-api-best-practices/internal/model"
	"github.com/credondocr/go-rest-api-best-practices/internal/service"
)

// GetAllProductsEndpoint godoc
// @Summary Get all products
// @Description Get all products using gRPC streaming
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} pb.Product
// @Router /products [get]
func GetAllProductsEndpoint(c *gin.Context) {
	fmt.Println("hols")
	products, err := service.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"message": "Error fetching products"})
		return
	}

	c.JSON(200, model.Response{Data: products})
}

// SearchProductsEndpoint godoc
// @Summary Get all products by fuzzy criteria
// @Description Get all products using gRPC streaming
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} pb.Product
// @Router /products/search [get]
func SearchProductsEndpoint(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		return
	}
	omitCache, _ := strconv.ParseBool(c.GetHeader("X-Omit-Cache"))
	products, err := service.SearchProducts(name, omitCache)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{Data: products, Metadata: model.Metadata{Count: len(products)}})
}
