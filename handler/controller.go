package handler

import (
	"fmt"
	"net/http"
	"rest-api_gin/transition"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama":   "author",
		"alamat": "nama alamat author",
	})
}

func Page2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "kosong",
	})
}

func UrlparamHandler(c *gin.Context) {
	id := c.Param("id")
	tahun := c.Param("tahun")
	c.JSON(http.StatusOK, gin.H{"url param data": id, "tahun": tahun})
}

func QueryparamHandler(c *gin.Context) {
	judul := c.Query("judul")
	rating := c.Query("rating")
	c.JSON(http.StatusOK, gin.H{"query param ? ": judul, "rating": rating})
}

func PostHandler(c *gin.Context) {
	var dataInput transition.ItemRequest
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
