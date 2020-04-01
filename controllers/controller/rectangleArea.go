package controller

import (
	interfaces "Start-Meli/controllers/Interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errRectangle error
var width float64
var height float64

//Rectangle  ...
type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Area ...
func (c Rectangle) Area() float64 {
	return c.Width * c.Height
}

// GETrectangleArea ...
func GETrectangleArea(c *gin.Context) {
	width, errRectangle = strconv.ParseFloat(c.Param("width"), 64)
	height, errRectangle = strconv.ParseFloat(c.Param("height"), 64)
	var rectangle = Rectangle{Width: width, Height: height}
	c.JSON(200, gin.H{"rectangle area": interfaces.Figura.Area(rectangle)})

}
