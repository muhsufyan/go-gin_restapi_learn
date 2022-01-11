// go mod init nama_projek
// install go-gin https://github.com/gin-gonic/gin
// go run main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// return statusnya yaitu 200, param 2 return objek interface jd bisa kembalikan apa saja (dlm hal ini kita kembalikan data)
		c.JSON(http.StatusOK, gin.H{
			"nama":   "author",
			"alamat": "nama alamat author",
		})
	})

	router.GET("/pageke2", func(c *gin.Context) {
		// return statusnya yaitu 200, param 2 return objek interface jd bisa kembalikan apa saja (dlm hal ini kita kembalikan data)
		c.JSON(http.StatusOK, gin.H{
			"data": "kosong",
		})
	})

	// liste/running
	router.Run(":8888")
}
