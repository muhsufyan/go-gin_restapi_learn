package handler

/*
kita akan update data dg id tertentu melalui http PUT.
data akan ditangkap oleh ItemRequest (request.go), disini kita gunakan ItemRequest karena data yg bisa di update itu sama dg create,
nah jika data diupdate beda maka buat saja DataRequestUpdate. tp dikasus ini kita tidak mengedit id sehingga sama sprti ItemRequest sblmnya yg dibuat untuk create
Kita harus buat func update yg di define dlm interface Repository (repository.go tambh func u/ update) lalu buat juga func update yg di define dlm interface Service (service.go tmbh func u/ update)
lalu buat controller untuk Update (kodenya dibawah dg func UpdateDataHandler), lalu buat path urlnya  di main.go
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

	dataRespone := convertToDataResponse(dataId)
	c.JSON(http.StatusOK, gin.H{
		"data": dataRespone,
	})
}

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
		"data": convertToDataResponse(datum),
	})
}

// Update
func (h *dataHandler) UpdateDataHandler(c *gin.Context) {

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
	// get id yg ingin diupdate
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	// Param 1 = id,
	datum, err := h.dataService.Update(id, dataRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToDataResponse(datum),
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
