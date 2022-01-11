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

	router.GET("/", rootHandler)
	router.GET("/pageke2", page2Handler)
	router.GET("/item/:id/:tahun", urlparamHandler)
	router.GET("/query", queryparamHandler)
	router.POST("item", postHandler)
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

// menangkap data yg diinput oleh user
type ItemInput struct {
	// data yg ditangkap akan disimpan disini
	Judul string `json:"judul" binding:"required"` //required hrs diisi
	//asalnya int agar tdk error ketika "500" maka dijdkan json.Number  tp ini masih error, sangat aneh
	Rating   json.Number `json:"rating" binding:"required,number"` //harus diisi dan berupa angka. gunakan playground untuk validasi
	SubTitle string      `json:"sub_title"`
}

// return data yg di input oleh user
func postHandler(c *gin.Context) {
	// define struct untuk tangkap dan isi
	var dataInput ItemInput

	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		// jika ada error maka server akan langsung mati, jd kode dibawh kita ganti
		// log.Fatal(err)

		// untuk tampung semua error
		errMsgs := []string{}
		// simpan semua error dan cetak
		for _, e := range err.(validator.ValidationErrors) {
			// cetak error
			errorMsg := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errMsgs = append(errMsgs, errorMsg)
			// handle error agar server tdk mati
			// c.JSON(http.StatusBadRequest, errorMsg)
			// fmt.Println(err)
		}
		// return semua error ke user
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
