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

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "17", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "53", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughn"},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
