package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func isLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := c.GetBool("is_logged_in")
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func isNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := c.GetBool("is_logged_in")
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// Set the "is_logged_in" variable in the request's context
// if the "login" cookie exists.
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
