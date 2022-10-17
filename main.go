package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Use(setUserStatus())

	initRoutes(router)

	router.Run(":80")
}
