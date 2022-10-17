package main

import (
	"demo-gin-app/handlers"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	// Index page
	router.GET("/", handlers.IndexPageHandler)

	// Login & Register user routes
	router.GET("/login", isNotLoggedIn(), handlers.LoginGetHandler)
	router.POST("/login", isNotLoggedIn(), handlers.LoginPostHandler)
	router.GET("/logout", isLoggedIn(), handlers.LogoutGetHandler)
	router.GET("/register", isNotLoggedIn(), handlers.RegisterGetHandler)
	router.POST("/register", isNotLoggedIn(), handlers.RegisterPostHandler)

	// Article routes
	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handlers.GetArticleHandler)
		articleRoutes.GET("/create", isLoggedIn(), handlers.CreateaArticleGetHandler)
		articleRoutes.POST("/create", isLoggedIn(), handlers.CreateaArticlePostHandler)
	}
}
