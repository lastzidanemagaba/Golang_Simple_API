package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CRUD struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()

	// Apply CORS middleware
	router.Use(corsMiddleware())

	// Define routes
	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/api/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	router.GET("/api/tutorials", func(c *gin.Context) {
		dataArray := []CRUD{
			{ID: 1, Title: "Title 1", Description: "Description 1"},
			{ID: 2, Title: "Title 2", Description: "Description 2"},
			{ID: 3, Title: "Title 3", Description: "Description 3"},
			{ID: 4, Title: "Title 4", Description: "Description 4"},
		}
		c.JSON(http.StatusOK, dataArray)
	})

	// Run the server
	router.Run(":3000")
}

func corsMiddleware() gin.HandlerFunc {
	var whiteList = []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, host := range whiteList {
			if host == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token, X-Auth-Email, X-Auth-Password, X-Auth-Name, X-Auth-Phone, X-Auth-Address, X-Auth-Role, X-Auth-Status, X-Auth-Image, X-Auth-Image-Name, X-Auth-Image-Path, X-Auth-Image-Size, X-Auth-Image-Type, X-Auth-Image-Ext, X-Auth-Image-Base64, X-Auth-Image-Base64-Name, X-Auth-Image-Base64-Path, X-Auth-Image-Base64-Size, X-Auth-Image-Base64-Type, X-Auth-Image-Base64-Ext, X-Auth-Image-Base64-Base64, X-Auth-Image-Base64-Base64-Name, X-Auth-Image-Base64-Base64-Path, X-Auth-Image-Base64-Base64-Size, X-Auth-Image-Base64-Base64-Type, X-Auth-Image-Base64-Base64-Ext, X-Auth-Image-Base64-Base64-Base64, X-Auth-Image-Base64-Base64-Base64-Name, X-Auth-Image-Base64-Base64-Base64-Path, X-Auth-Image-Base64-Base64-Base64-Size, X-Auth-Image-Base64-Base64-Base64-Type, X-Auth-Image-Base64-Base64-Base64-Ext, X-Auth-Image-Base64-Base64-Base64-Base64, X-Auth-Image-Base64-Base64-Base64-Base64-Name, X-Auth-Image-Base64-Base64-Base64-Base64-Path, X-Auth-Image-Base64-Base64-Base64-Base64-Size, X-Auth-Image-Base64-Base64-Base64-Base64-Type, X-Auth-Image-Base64-Base64-Base64-Base64-Ext, X-Auth-Image-Base64-Base64-Base64-Base")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
