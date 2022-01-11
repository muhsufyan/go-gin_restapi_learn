package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/pageke2", page2Handler)
	// url parameter
	router.GET("/item/:id", urlparamHandler)
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
	// tangkap url parameter, dlm hal ini id, dan simpan kedlm variabel id
	id := c.Param("id")

	// kirim id tsb sbg respon (data)ke user
	c.JSON(http.StatusOK, gin.H{"url param data": id})
}
