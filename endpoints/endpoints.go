package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	albums, err := LoadAlbumsFromFile("albums.json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error loading albums"})
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsById(c *gin.Context) {
	albums, err := LoadAlbumsFromFile("albums.json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error loading albums"})
	}

	id := c.Param("id")

	for index, value := range albums {
		if albums[index].ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetAlbumsByIdQuery(c *gin.Context) {
	albums, err := LoadAlbumsFromFile("albums.json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error loading albums"})
	}

	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusOK, albums)
		return
	}

	for index, value := range albums {
		if albums[index].ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	// Package level variable
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Load the albums from file, expose error to caller (unsafe technically)
	albums, err := LoadAlbumsFromFile("albums.json")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error loading albums: %v", err)})
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	// Save the updated slice to file, expose error to caller (unsafe technically)
	err = SaveAlbumsToFile("albums.json", albums)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error saving albums: %v", err)})
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
