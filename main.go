package main

import (
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

	//we can apply if else jd if x maka simpan ke db, if else y maka simpan ke file text

	// simpan ke db
	// dataRepository := transition.NewRepository(db)
	// dataService := transition.NewService(dataRepository)//simpan ke db(dilakukan melalui param dataRepository yg mrpkn interface dr Repository)
	// dataHandler := handler.NewDataHandler(dataService)

	// pura"nya simpan ke file text
	// dataRepository := transition.NewRepository(db)
	// instansiasi untuk fileRepository
	dataFileRepository := transition.NewFileRepository()
	dataService := transition.NewService(dataFileRepository) // lewat paramnya kita simpan data kedlm file text melalui interface fileRepository
	dataHandler := handler.NewDataHandler(dataService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", dataHandler.RootHandler)
	v1.GET("/pageke2", dataHandler.Page2Handler)
	v1.GET("/item/:id/:tahun", dataHandler.UrlparamHandler)
	v1.GET("/query", dataHandler.QueryparamHandler)
	v1.POST("/item", dataHandler.PostHandler)

	router.Run(":8888")
}
