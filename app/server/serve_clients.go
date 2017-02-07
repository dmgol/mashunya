package server

import (
	"fmt"
	"net/http"

	"github.com/dmgol/mashunya/app/models"
	"github.com/dmgol/mashunya/config"
	"github.com/dmgol/mashunya/db"
	"github.com/dmgol/mashunya/utils"
	"github.com/gin-gonic/gin"
)

var resultNotFound = gin.H{
	"Result": "Not found",
}

func getCollectionList(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, resultNotFound)
	return
}

func selectProductsByCollection(collectionID uint, filter models.Product) (statusCode int, result interface{}) {
	var collection models.Collection
	if db.DB.Debug().Preload("Products", filter).First(&collection, collectionID).RecordNotFound() {
		return http.StatusNotFound, resultNotFound
	}
	return http.StatusOK, models.ProductList(collection.Products).H()
}

func selectProducts(filter models.Product) (statusCode int, result interface{}) {
	var products []models.Product

	if db.DB.Debug().Limit(100).Preload("ColorVariations").Preload("ColorVariations.Color").Preload("ColorVariations.SizeVariations.Size").Find(&products, filter).RecordNotFound() {
		return http.StatusNotFound, resultNotFound
	}

	return http.StatusOK, models.ProductList(products).H()
}

func getProductList(ctx *gin.Context) {
	var filter models.Product

	utils.ParseUintRef(ctx.Query("category"), &filter.CategoryID)

	if collectionID, success := utils.ParseUint(ctx.Query("collection")); success {
		ctx.JSON(selectProductsByCollection(uint(collectionID), filter))
	} else {
		ctx.JSON(selectProducts(filter))
	}

}

func selectProduct(code string) (statusCode int, result interface{}) {
	var product models.Product
	if db.DB.Debug().Preload("ColorVariations").Preload("ColorVariations.Color").Preload("ColorVariations.SizeVariations.Size").Where(&models.Product{Code: code}).First(&product).RecordNotFound() {
		return http.StatusNotFound, resultNotFound
	}
	return http.StatusOK, product.H()
}

func getProduct(ctx *gin.Context) {
	productCode := ctx.Param("code")
	ctx.JSON(selectProduct(productCode))
}

func serveStaticPath(router *gin.Engine, paths []string, root string) {
	for _, path := range paths {
		router.Static("/"+path, "public/"+path+"/"+root)
	}
}

// ServeClients serves client request
func ServeClients() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Result": "Ok"})
	})
	router.GET("/products", getProductList)
	router.GET("/products/:code", getProduct)

	serveStaticPath(router, []string{"system"}, config.Root)

	router.Run(fmt.Sprintf(":%d", config.Config.ClientsPort))
}
