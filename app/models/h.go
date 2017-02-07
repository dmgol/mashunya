package models

import "github.com/gin-gonic/gin"

type ProductList []Product

func (list ProductList) H() []gin.H {
	result := make([]gin.H, len(list))
	for i, el := range list {
		result[i] = el.H()
	}
	return result
}

func (p Product) H() gin.H {
	return gin.H{
		"Name":            p.Name,
		"Code":            p.Code,
		"Price":           p.Price,
		"ColorVariations": ColorVariationList(p.ColorVariations).H(),
		"Path":            p.DefaultPath(),
		"MainImageUrl":    p.MainImageURL(),
	}
}

type ColorVariationList []ColorVariation

func (list ColorVariationList) H() []gin.H {
	result := make([]gin.H, len(list))
	for i, el := range list {
		result[i] = el.H()
	}
	return result
}

func (cv ColorVariation) H() gin.H {
	return gin.H{
		"Color":          cv.Color.H(),
		"SizeVariations": SizeVariationList(cv.SizeVariations).H(),
	}
}

type SizeVariationList []SizeVariation

func (list SizeVariationList) H() []gin.H {
	result := make([]gin.H, len(list))
	for i, el := range list {
		result[i] = el.H()
	}
	return result
}

func (sv SizeVariation) H() gin.H {
	return gin.H{
		"Size": sv.Size.H(),
	}
}

func (size Size) H() gin.H {
	return gin.H{
		"Name": size.Name,
	}
}

func (color Color) H() gin.H {
	return gin.H{
		"Name": color.Name,
	}
}
