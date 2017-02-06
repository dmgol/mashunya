package models

import "github.com/gin-gonic/gin"

type H interface {
	H() gin.H
}

func iterate(array []H) []gin.H {
	result := make([]gin.H, len(array))
	for i, el := range array {
		result[i] = el.H()
	}
	return result
}

func (p Product) H() gin.H {
	colorVars := make([]gin.H, len(p.ColorVariations))

	for i, v := range p.ColorVariations {
		colorVars[i] = v.H()
	}

	return gin.H{
		"Name":            p.Name,
		"Code":            p.Code,
		"Price":           p.Price,
		"ColorVariations": colorVars,
	}
}

func (cv ColorVariation) H() gin.H {
	sizeVars := make([]gin.H, len(cv.SizeVariations))
	for i, sv := range cv.SizeVariations {
		sizeVars[i] = sv.H()
	}
	return gin.H{
		"Color":          cv.Color.H(),
		"SizeVariations": sizeVars,
	}
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
