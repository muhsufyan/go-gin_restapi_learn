package handler

import (
	"fmt"
	"net/http"
	"rest-api_gin/transition"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type dataHandler struct {
	dataService transition.Service
}

func NewDataHandler(dataService transition.Service) *dataHandler {

	return &dataHandler{dataService}
}

func (h *dataHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama":   "author",
		"alamat": "nama alamat author",
	})
}

func (h *dataHandler) Page2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "kosong",
	})
}

func (h *dataHandler) UrlparamHandler(c *gin.Context) {
	id := c.Param("id")
	tahun := c.Param("tahun")
	c.JSON(http.StatusOK, gin.H{"url param data": id, "tahun": tahun})
}

func (h *dataHandler) QueryparamHandler(c *gin.Context) {
	judul := c.Query("judul")
	rating := c.Query("rating")
	c.JSON(http.StatusOK, gin.H{"query param ? ": judul, "rating": rating})
}

func (h *dataHandler) PostHandler(c *gin.Context) {

	var dataRequest transition.ItemRequest

	err := c.ShouldBindJSON(&dataRequest)
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

	datum, err := h.dataService.Create(dataRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": datum,
	})
}
