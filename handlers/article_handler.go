package handlers

import (
	"demo-gin-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticleHandler(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
			c.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func CreateaArticleGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create-article.html", gin.H{
		"title": "Create New Article",
	})
}

func CreateaArticlePostHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if newArticle, err := models.CreateNewArticle(title, content); err == nil {
		c.HTML(http.StatusOK, "article-submission-successful.html", gin.H{
			"title":   "Successful article submission",
			"payload": newArticle,
		})
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
