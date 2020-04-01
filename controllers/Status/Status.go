package status

import (
	"github.com/gin-gonic/gin"
)

type status struct{}

/*
Ping return a message pong
*/
func Ping(c *gin.Context) {
	h := status{}
	h.ping(c)
}

func (h status) ping(c *gin.Context) {
	//return message "pong" and status code 200
	c.JSON(200, gin.H{
		"message": "pongg",
	})
}
