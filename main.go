// install gorm dan install driver mysql jika error
package main

import (
	"fmt"
	"log"
	"rest-api_gin/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// koneksi ke db
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/simpan?charset=utf8mb4&parseTime=True&loc=Local" //usernamenya root, passwordnya tdk ada, nama dbnya "simpan"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//cek jika ada error saat koneksi
	if err != nil {
		log.Fatal("koneksi ke db error")
	}
	fmt.Println("Database telah terkoneksi")
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("item", handler.PostHandler)

	router.Run(":8888")
}
