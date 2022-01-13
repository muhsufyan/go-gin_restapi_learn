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

	// UPDATE. untuk melakukan update kita perlu data apa yg ingin di update (dlm hal ini adlh id)
	var data transition.Penyimpanan

	// get id
	err = db.Debug().Where("id = ?", 1).First(&data).Error //SELECT * FROM `penyimpanans` WHERE id = 1 ORDER BY `penyimpanans`.`id` LIMIT 1
	if err != nil {
		fmt.Println("id tidak ditemukan")
	}

	// update judul dari id yg ditentukan
	data.Judul = "judul berhasil di update"
	// simpan ke db, jika tdk Save() maka data hanya disimpan dlm memory tdk tersimpan ke db
	err = db.Save(&data).Error
	if err != nil {
		fmt.Println("data tidak terupdate")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("/item", handler.PostHandler)

	router.Run(":8888")
}
