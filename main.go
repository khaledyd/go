package main

import (
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, albums)
	})

	r.POST("/", func(c *gin.Context) {
		var newAlbum album

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		// Add the new album to the slice.
		albums = append(albums, newAlbum)
		c.JSON(200, albums)
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, a := range albums {
			if a.ID == id {
				albums = append(albums[:i], albums[i+1:]...)
				c.JSON(200, albums)
				return
			}
		}
		c.JSON(404, gin.H{"error": "album not found"})
	})

	r.PUT("/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedAlbum album

		if err := c.BindJSON(&updatedAlbum); err != nil {
			return
		}

		for i, a := range albums {
			if a.ID == id {
				albums[i] = updatedAlbum
				c.JSON(200, albums)
				return
			}
		}
		c.JSON(404, gin.H{"error": "album not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
