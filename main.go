// kasusnya misal data judul parameternya diganti jd title sedangkan client masih ada yg memakai parameter yg lama yaitu judul
// sehingga parameter yg lama (judul) hrs tetap ada u/ mempertahankan client yg blm update
// solusinya dg membuat group (grouping router) sehingga ada versi 1(v1), v2, vn
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	// grouping router
	v1 := router.Group("/v1")
	// kita jdkan router dibawah termasuk group v1. misal akses item jd .../v1/item
	v1.GET("/", rootHandler)
	v1.GET("/pageke2", page2Handler)
	v1.GET("/item/:id/:tahun", urlparamHandler)
	v1.GET("/query", queryparamHandler)
	v1.POST("item", postHandler)

	// versioning 2
	v2 := router.Group("/v2")
	v2.GET("/", rootHandler)

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

func queryparamHandler(c *gin.Context) {
	judul := c.Query("judul")
	rating := c.Query("rating")
	c.JSON(http.StatusOK, gin.H{"query param ? ": judul, "rating": rating})
}

type ItemInput struct {
	Judul    string      `json:"judul" binding:"required"`
	Rating   json.Number `json:"rating" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
}

func postHandler(c *gin.Context) {
	var dataInput ItemInput
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		errMsgs := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errMsgs = append(errMsgs, errorMsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMsgs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"judul":     dataInput.Judul,
		"rating":    dataInput.Rating,
		"sub_title": dataInput.SubTitle,
	})
}
