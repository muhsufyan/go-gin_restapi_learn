package handler

import (
	"fmt"
	"net/http"
	"rest-api_gin/transition"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// buat struct handler untuk return di NewHandler
type dataHandler struct {
	// bth service
	dataService transition.Service
}

//paramnya interface Service
func NewDataHandler(dataService transition.Service) *dataHandler {
	// passing dataService
	return &dataHandler{dataService}
}

// agar semua function yg mengatur link/ url sprti RootHandler, PostHandler, dll dpt  dimiliki oleh struct dataHandler maka ubah function jd jd method.
// contoh dr function func RootHandler(c *gin.Context) {}
// contoh function yg dimiliki oleh sebuah struct func (handler *dataHandler) RootHandler(c *gin.Context) {} dan ini disbt method
// jd kita tambh semua function dg (handler *dataHandler)
// lalu handler itu punya service yaitu dataService sehingga didlm setiap method dpt memanggil servie (dataService) tp kita hrs buat dulu dataHandler nya di main
// agar lbh mudah kita singkat handler *dataHandler jd h *dataHandler, supaya konsisten sblmnya juga service *service jd s *service
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
	// dataRequest artinya data yg berasal dr user yg meminta u/ disimpan kedlm db
	var dataRequest transition.ItemRequest
	// data ditangkap disini. di binding disini
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
	// panggil servicenya
	datum, err := h.dataService.Create(dataRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// balikin data dlm var datum
	c.JSON(http.StatusOK, gin.H{
		"data": datum,
	})
}
