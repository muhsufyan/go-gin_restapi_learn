package transition

import "encoding/json"

type ItemInput struct {
	Judul    string      `json:"judul" binding:"required"`
	Rating   json.Number `json:"rating" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
}
