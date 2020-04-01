package main

import (
	"Start-Meli/controllers/controller"
	"Start-Meli/controllers/status"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", status.Ping)
	r.GET("/circlearea/:radius", controller.GETCircleArea)
	r.GET("/rectanglearea/:width/:height", controller.GETrectangleArea)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
