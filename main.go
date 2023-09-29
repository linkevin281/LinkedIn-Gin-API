package main

import (
	"GinTesting/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	// Function level variable
	router := gin.Default()
	router.POST("/albums", endpoints.PostAlbums)
	router.GET("/albums/:id", endpoints.GetAlbumsById)
	router.GET("/albums", endpoints.GetAlbumsByIdQuery)
	router.Run("localhost:8080")
}
