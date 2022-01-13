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
	// sblmnya kita melakukan query dg objek db scra langsung. now kita akan bungkus objek db kedlm repository
	// pertama kita buat repository simpan yg bertanggung jwb terhdp entity penyimpanan (tabel penyimpanans)
	// dlm transition buat file baru repository.go
	// dg repository kita membatasi agar mengakses db tdk dilakukan scra langsung
	// now ketika melakukan query kita bikin struct lalu buat function dulu sesuai dg tugas querynya lalu panggil dbnya dlm repository
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("koneksi ke db error")
	}

	db.AutoMigrate(&transition.Penyimpanan{})

	// COBA TAMPILKAN SEMUA JUDUL
	// passing objek db ke NewRepository
	dataRespository := transition.NewRepository(db)
	// ambil semua data di Penyimpanan
	dataset, err := dataRespository.FindAll()
	if err != nil {
		fmt.Println("data tidak ditemukan ERROR")
	}

	// fetch datanya
	for _, data := range dataset {
		fmt.Println("Judul :", data.Judul)
	}

	// GET DATA BY ID TERTENTU
	data, err := dataRespository.FindById(3)
	fmt.Println("Judul dari id ke 3 adlh", data.Judul)

	// CREATE DATA BARU
	dataIn := transition.Penyimpanan{
		Judul:    "judul baru dari repository",
		Rating:   4,
		SubTitle: "en,in",
	}
	// simpan
	dataRespository.Create(dataIn)
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("/item", handler.PostHandler)

	router.Run(":8888")
}
