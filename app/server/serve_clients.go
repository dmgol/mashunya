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

func getCollection(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, resultNotFound)
	return
}

func getProductList(ctx *gin.Context) {
	var (
		products []models.Product
		filter   models.Product
	)

	categoryID, err := strconv.ParseUint(ctx.Query("category"), 10, 32)
	if err == nil {
		filter.CategoryID = uint(categoryID)
	}

	collectionID, err := strconv.ParseUint(ctx.Query("collection"), 10, 32)
	if err == nil {
		var collection models.Collection
		if db.DB.Debug().Preload("Products", filter).First(&collection, collectionID).RecordNotFound() {
			ctx.JSON(http.StatusNotFound, resultNotFound)
			return
		}
		ctx.JSON(http.StatusOK,
			gin.H{
				"Result": collection.Products,
			})
		return
	}

	if db.DB.Debug().Limit(100).Find(&products, filter).RecordNotFound() {
		ctx.JSON(http.StatusNotFound, resultNotFound)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"Result": products,
		})
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
