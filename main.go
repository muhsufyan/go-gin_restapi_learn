// install gorm dan install driver mysql jika error
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
	// koneksi ke db
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:@tcp(127.0.0.1:3306)/simpan?charset=utf8mb4&parseTime=True&loc=Local" //usernamenya root, passwordnya tdk ada, nama dbnya "simpan"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//cek jika ada error saat koneksi
	if err != nil {
		log.Fatal("koneksi ke db error")
	}

	// migrate agar tidak perlu buat tabel didb tp cukup disini(kasus ini ada di struct penyimpanan filenya transition/entity), sprti migrasinya laravel
	// akan dibuat tabel dg nama penyimpanans karena akan berbentuk plural
	db.AutoMigrate(&transition.Penyimpanan{})
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/pageke2", handler.Page2Handler)
	v1.GET("/item/:id/:tahun", handler.UrlparamHandler)
	v1.GET("/query", handler.QueryparamHandler)
	v1.POST("item", handler.PostHandler)

	router.Run(":8888")
}
