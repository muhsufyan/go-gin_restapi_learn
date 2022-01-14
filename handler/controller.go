package handler

/*
get data tp key jsonnya huruf kecil semua, dan hanya menampilkan data yg diinginkan saja (pd bagian get all data akan kembalikan semua data lewat struct Penyimpanan di entity.go)
untuk melakukan itu kita hrs buat struct yg akan mewakili data json response, dlm transition buat yaitu response.go
*/
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

func (h *dataHandler) GetDataset(c *gin.Context) {
	// kode dibawah ini akan mengarah ke struct Penyimpanan jd kita ubah sehingga mengarah ke struct dataResponse
	dataset, err := h.dataService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// disini ubah arah jd ke struct dataResponse
	var datasetResponse []transition.DataResponse
	// mapping datanya
	for _, datum := range dataset {
		dataResponse := transition.DataResponse{
			ID:       datum.ID,
			Judul:    datum.Judul,
			Rating:   datum.Rating,
			SubTitle: datum.SubTitle,
		}
		datasetResponse = append(datasetResponse, dataResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": datasetResponse,
	})
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
