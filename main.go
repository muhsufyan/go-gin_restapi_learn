package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/pageke2", page2Handler)
	router.GET("/item/:id/:tahun", urlparamHandler)
	// query param handler, misal /drama?judul=data_input. ini akan return data_input
	router.GET("/query", queryparamHandler)
	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama":   "author",
		"alamat": "nama alamat author",
	})
}

func page2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "kosong",
	})
}

func urlparamHandler(c *gin.Context) {
	id := c.Param("id")
	tahun := c.Param("tahun")
	c.JSON(http.StatusOK, gin.H{"url param data": id, "tahun": tahun})
}

// cara kerjanya mirip sprti urlparamHandler
func queryparamHandler(c *gin.Context) {
	judul := c.Query("judul")
	rating := c.Query("rating")
	c.JSON(http.StatusOK, gin.H{"query param ? ": judul, "rating": rating})
}
