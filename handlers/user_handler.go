package handlers

import (
	"demo-gin-app/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func LoginPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if models.IsUserValid(username, password) {
		// Set the "session" token.
		token := strconv.FormatInt(rand.Int63(), 16)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		c.HTML(http.StatusOK, "login-successful.html", gin.H{
			"title": "Successful Login",
		})
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error_title":   "Login failed",
			"error_message": "Invalid credentials provided",
		})
	}
}

func LogoutGetHandler(c *gin.Context) {
	// Set empty string for login/logout token.
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the homepage.
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func RegisterGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func RegisterPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// Set the "session" token.
		token := strconv.FormatInt(rand.Int63(), 16)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		c.HTML(http.StatusOK, "registration-successful.html", gin.H{
			"title": "Successful Registration & Login",
		})
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error_title":   "Registration failed",
			"error_message": err.Error(),
		})
	}
}
