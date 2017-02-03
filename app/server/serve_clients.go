package server

import (
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"github.com/dmgol/mashunya/app/models"
	"github.com/dmgol/mashunya/config"
	"github.com/dmgol/mashunya/db"
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
	return http.StatusOK, collection.Products
}

func selectProducts(filter models.Product) (statusCode int, result interface{}) {
	var products []models.Product

	if db.DB.Debug().Limit(100).Find(&products, filter).RecordNotFound() {
		return http.StatusNotFound, resultNotFound
	}

	return http.StatusOK, products
}

func str2uint(str string, out *uint) bool {
	if value, err := strconv.ParseUint(str, 10, 32); err == nil {
		*out = uint(value)
		return true
	}
	return false
}

func str2uintResult(p string) (value uint, success bool) {
	success = str2uint(p, &value)
	return
}

func getProductList(ctx *gin.Context) {
	var filter models.Product

	str2uint(ctx.Query("category"), &filter.CategoryID)

	if collectionID, success := str2uintResult(ctx.Query("collection")); success {
		ctx.JSON(selectProductsByCollection(uint(collectionID), filter))
	} else {
		ctx.JSON(selectProducts(filter))
	}

}

func getProduct(ctx *gin.Context) {
	var (
		product models.Product
		//colorVariation models.ColorVariation
		codes       = strings.Split(ctx.Param("code"), "_")
		productCode = codes[0]
		//colorCode      string
	)

	// if len(codes) > 1 {
	// 	colorCode = codes[1]
	// }

	if db.DB.Debug().Preload("ColorVariations").Preload("ColorVariations.Color").Preload("Category").Where(&models.Product{Code: productCode}).First(&product).RecordNotFound() {
		ctx.JSON(http.StatusNotFound,
			gin.H{
				"Product": "Not found",
			})
		return
	}

	//db.DB.Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)

	ctx.JSON(http.StatusOK,
		gin.H{
			"Product": product,
			//"ColorVariation": colorVariation,
		})
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
