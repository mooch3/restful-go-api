package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	// serialize the struct into JSON and add it to the response
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _,a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no album found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	// call BindJSON to bind the received JSON to the newAlbum variable
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "17", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "53", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughn"},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
