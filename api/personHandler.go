package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add person details Add Handler
func AddPersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewList()
		c.Bind(&requestBody)
		AddPersonDetailsService(requestBody)
		c.JSON(http.StatusOK, GetPersonList())
	}
}

// Update person details Handler
func UpdatePersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewList()
		c.Bind(&requestBody)
		UpdatePersonDetailsService(requestBody)
		c.JSON(http.StatusOK, GetPersonList())
	}
}

// Delete person details Delete Handler
func DeletePersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewList()
		c.Bind(&requestBody)
		DeletePersonDetailsService(requestBody)
		c.JSON(http.StatusOK, GetPersonList())
	}
}

// View person details Handler
func ViewPersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, GetPersonList())
	}
}
