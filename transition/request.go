package transition

import (
	"encoding/json"
)

// type ItemInput struct {
// code diatas diganti jd
type ItemRequest struct {
	Judul    string      `json:"judul" binding:"required"`
	Rating   json.Number `json:"rating" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
}
