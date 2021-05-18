package main

import (
	"test3/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/createCode", controller.Create)

	r.GET("/inquire", controller.Inquire)

	r.GET("/client", controller.Client)


	r.Run()
}
