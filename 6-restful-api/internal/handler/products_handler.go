package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/domain"
	"github.com/muhangga/internal/service"
	"github.com/muhangga/internal/utils"
)

type productHandler struct {
	productsService service.ProductsService
}

func NewProductHandler(productsService service.ProductsService) *productHandler {
	return &productHandler{productsService}
}

func (h *productHandler) FetchAllProducts(c *gin.Context) {
	products, err := h.productsService.FetchAllProducts()
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := utils.BuildResponse(true, "Successfully fetch products", products)
	c.JSON(http.StatusOK, resp)
}

func (h *productHandler) CreateProduct(c *gin.Context) {

	var productRequest domain.ProductRequest
	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	product, err := h.productsService.CreateProduct(productRequest)
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.BuildResponse(true, "Products successfully created", product)
	c.JSON(http.StatusOK, resp)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var productRequest domain.ProductRequest
	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productsService.UpdateProduct(id, productRequest)
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.BuildResponse(true, "Products successfully updated", product)
	c.JSON(http.StatusOK, resp)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.productsService.DeleteProduct(id)
	if err != nil {
		resp := utils.BuildResponse(false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.BuildResponse(true, "Products successfully deleted", true)
	c.JSON(http.StatusOK, resp)
}
