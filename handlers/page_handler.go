package handlers

import (
	"demo-gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	isLoggedIn := c.GetBool("is_logged_in")
	articles := models.GetAllArticles()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":        "Home Page",
		"is_logged_in": isLoggedIn,
		"payload":      articles,
	})
}
