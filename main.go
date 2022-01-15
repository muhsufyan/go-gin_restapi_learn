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

	dataRepository := transition.NewRepository(db)
	dataService := transition.NewService(dataRepository)
	dataHandler := handler.NewDataHandler(dataService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", dataHandler.GetDataset)
	v1.GET("/getById/:id", dataHandler.GetDataByIdHandler)
	v1.POST("/item", dataHandler.CreateNewDataHandler)
	v1.PUT("/update/:id", dataHandler.UpdateDataHandler)
	v1.DELETE("/delete/:id", dataHandler.DeleteDataHandler)

	router.Run(":8888")
}
