package handler

/* endpoint untuk get all data
service dan repository sdh selesai (lwt func FindAll), tinggal dibagian handler/controller ini dan kita hanya gunakan PostHandler jd handler yg lainnya dihapus saja
kita buat endpoint baru di main.go
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

// get semua data, kita sebut semua data as dataset
func (h *dataHandler) GetDataset(c *gin.Context) {
	dataset, err := h.dataService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// jika ada semua data, kita return
	c.JSON(http.StatusOK, gin.H{
		"data": dataset,
	})
}

// create data
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
