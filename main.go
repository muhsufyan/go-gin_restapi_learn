package main

// now kita input json lwt postman http post v1/item yg akan disimpan kedlm db
// mengakses service dr dlm handler(how handler access service), caranya sama saja dg mengakses repository dr dlm service/how service can access repository
// dlm handler buat func baru New....
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

	// buat dataHandler, passing (lwt param) fungsi dataServicenya
	dataHandler := handler.NewDataHandler(dataService)
	router := gin.Default()

	v1 := router.Group("/v1")
	// cara akses nya diubah jd lwt Service dulu tdk langsung ke handler
	v1.GET("/", dataHandler.RootHandler)
	v1.GET("/pageke2", dataHandler.Page2Handler)
	v1.GET("/item/:id/:tahun", dataHandler.UrlparamHandler)
	v1.GET("/query", dataHandler.QueryparamHandler)
	v1.POST("/item", dataHandler.PostHandler)

	router.Run(":8888")
}
