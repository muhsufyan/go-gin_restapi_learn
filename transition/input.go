package transition

import "encoding/json"

// struct itu lbh ke arah transisi dari model ke logik/handler/controller. jd kita simpan kedlm folder transition
type ItemInput struct {
	Judul    string      `json:"judul" binding:"required"`
	Rating   json.Number `json:"rating" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
}
