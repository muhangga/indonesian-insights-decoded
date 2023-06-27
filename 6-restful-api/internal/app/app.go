package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/handler"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/service"
)

func NewInitializedServer(cfg config.Config) *gin.Engine {
	router := gin.Default()

	db := config.NewDB(cfg)

	configureCors(router)

	api := router.Group("/api")

	productRepo := repository.NewProductsRepository(db)
	productService := service.NewProductsService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	api.GET("/products", productHandler.FetchAllProducts)
	api.POST("/products", productHandler.CreateProduct)
	api.PUT("/product/:id", productHandler.UpdateProduct)
	api.DELETE("/product/:id", productHandler.DeleteProduct)

	return router
}

func configureCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))
}
