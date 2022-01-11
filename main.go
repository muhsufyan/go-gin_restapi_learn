package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/pageke2", page2Handler)

	router.Run(":8888")
}

// param nya berupa pointer yg mengarah ke gin.Context disimpan kedlm c
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
