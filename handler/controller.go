package handler

/*
get data dg id tertentu
localhost/v1/2 , disana 2 adlh id jd yg dikembalikan hanya data dg id ke 2 saja.
Repository dan Service nya sdh ada tinggal handler saja
*/
import (
	"fmt"
	"net/http"
	"rest-api_gin/transition"
	"strconv"

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

	dataset, err := h.dataService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var datasetResponse []transition.DataResponse
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

// get data id tertentu
func (h *dataHandler) GetDataByIdHandler(c *gin.Context) {
	// tangkap data id nya, id berapa ?
	idStr := c.Param("id")
	// convert id dr string ke int
	id, _ := strconv.Atoi(idStr)
	dataId, err := h.dataService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// return ke user semua data termasuk CreateAt dan UpdateAt
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": dataId,
	// })
	// kode diatas kita ubah sehingga hanya mengembalikan data tanpa CreateAt dan UpdateAt, sama sprti sblmnya
	dataRespone := transition.DataResponse{
		ID:       dataId.ID,
		Judul:    dataId.Judul,
		Rating:   dataId.Rating,
		SubTitle: dataId.SubTitle,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": dataRespone,
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
