package handler

/*
sblmnya kode untuk return ke user dg data tertentu hrs dilakukan berulang(tdk reuse), kita atasi mslh tsb dg melakukan refaktor/ dlm kasus
ini disbt response converter. kode yg akan di refaktor tsb adlh
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

dan kode dibwh ini juga sama
dataId, err := h.dataService.FindByID(int(id))
dataRespone := transition.DataResponse{
		ID:       dataId.ID,
		Judul:    dataId.Judul,
		Rating:   dataId.Rating,
		SubTitle: dataId.SubTitle,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": dataRespone,
	})

terlihatkan tdk reuse saat mapping data.
caranya dg membuat private func (lihat dipaling bwh)
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
		// mapping data jd sesuai yg diinginkan tinggal panggil convertToDataResponse dg param datanya
		dataResponse := convertToDataResponse(datum)
		datasetResponse = append(datasetResponse, dataResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": datasetResponse,
	})
}

func (h *dataHandler) GetDataByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	dataId, err := h.dataService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// mapping data jd sesuai yg diinginkan tinggal panggil convertToDataResponse dg param datanya
	dataRespone := convertToDataResponse(dataId)
	c.JSON(http.StatusOK, gin.H{
		"data": dataRespone,
	})
}

// nama func PostHandler diubah jd CreateNewDataHandler
func (h *dataHandler) CreateNewDataHandler(c *gin.Context) {

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

func convertToDataResponse(dataObj transition.Penyimpanan) transition.DataResponse {
	return transition.DataResponse{
		ID:       dataObj.ID,
		Judul:    dataObj.Judul,
		Rating:   dataObj.Rating,
		SubTitle: dataObj.SubTitle,
	}
}
