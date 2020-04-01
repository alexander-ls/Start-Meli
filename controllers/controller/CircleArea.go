package controller

import (
	interfaces "Start-Meli/controllers/Interfaces"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error
var radius float64

//Circle ...
type Circle struct {
	Radius float64 `json:"radius"`
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// GETCircleArea ...
func GETCircleArea(c *gin.Context) {
	radius, err = strconv.ParseFloat(c.Param("radius"), 64)
	var circle = Circle{Radius: radius}
	fmt.Println(Circle.Area(circle))
	c.JSON(200, gin.H{"Circle's area": interfaces.Figura.Area(circle)})

}
