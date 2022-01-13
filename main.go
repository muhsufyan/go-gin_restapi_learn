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

	// instansiasi repository
	dataRepository := transition.NewRepository(db)
	// buat instansiasi service, paramnya adlh interface Repository jd kita perlu struct yg mengimplement interface Repository
	dataService := transition.NewService(dataRepository)

	// ini seharusnya ada dihandler /controller lbh tptnya di func PostHandler, tp hanya percobaan jd disini dulu saja
	dataRequest := transition.ItemRequest{
		Judul:    "service & repository",
		Rating:   "200",
		SubTitle: "melayu",
	}
	dataService.Create(dataRequest)
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("/item", handler.PostHandler)

	router.Run(":8888")
}
