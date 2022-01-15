package handler

/*
delete ini hampir sama dg update (alur)
1. buat deklarasi func delete di interface Repository (agar data dpt terhapus di db) dan buat implement-nya. pergi ke repostiory.go
2. buat juga di service. di service perlu menerima data id untuk data yg ingin dihapus. pergi ke service.go
3. buat DeleteDataHandler untuk dipanggil di main.go nantinya
4. buat http delete di main.go yg memanggil DeleteDataHandler
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

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
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

// DELETE HANDLER/CONTROLLER
func (h *dataHandler) DeleteDataHandler(c *gin.Context) {
	// tangkap id yg ingin dihapus dari user
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	// jlnkan Service delete
	dataId, err := h.dataService.Delete(int(id))
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

func convertToDataResponse(dataObj transition.Penyimpanan) transition.DataResponse {
	return transition.DataResponse{
		ID:       dataObj.ID,
		Judul:    dataObj.Judul,
		Rating:   dataObj.Rating,
		SubTitle: dataObj.SubTitle,
	}
}
