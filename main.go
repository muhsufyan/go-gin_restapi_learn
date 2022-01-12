package main

import (
	"fmt"
	"log"
	"rest-api_gin/handler"
	"rest-api_gin/transition"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/simpan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("koneksi ke db error")
	}

	db.AutoMigrate(&transition.Penyimpanan{})

	// create data baru ke tabel penyimpanans(lwt struct Penyimpanan yg ada di transition/entity)
	// buat objek dr struct Penyimpanan
	data := transition.Penyimpanan{}
	// data.ID sudah digenerate o/ db (autoincrement)
	data.Judul = "sword"
	data.Rating = 3
	data.SubTitle = "en"
	// data.CreatedAt & data.UpdatedAt dibuat otomatis
	// simpan ke db sebagai data baru
	// db.Create(&data) kenapa pakai tanda & karena as pointer
	// cek dulu jika ada error
	err = db.Create(&data).Error
	if err != nil {
		fmt.Println("Error saat menyimpan data baru")
	}
	// ini hanya sementara jd blm selesai sampai sini
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("/item", handler.PostHandler)

	router.Run(":8888")
}
