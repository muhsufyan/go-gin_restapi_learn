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

	// // create new data
	// data := transition.Penyimpanan{}

	// data.Judul = "sword season 2"
	// data.Rating = 5
	// data.SubTitle = "in"

	// err = db.Create(&data).Error
	// if err != nil {
	// 	fmt.Println("Error saat menyimpan data baru")
	// }

	// dpt data
	var data transition.Penyimpanan

	// get 1 data pertama(di urutan pertama, order by id ascending), passing datanyadg menjdkan param
	// err = db.First(&data).Error
	// kode diatas dpt dimunculkan query sqlnya dg menambahkan Debug(). Jika ingin mengambil 1 data terakhir ganti First jd Last
	// err = db.Debug().First(&data).Error // query sqlnya : SELECT * FROM `penyimpanans` ORDER BY `penyimpanans`.`id` LIMIT 1
	// cari 1 data dg id tertentu
	err = db.Debug().First(&data, 2).Error // query sqlnya : SELECT * FROM `penyimpanans` WHERE `penyimpanans`.`id` = 2 ORDER BY `penyimpanans`.`id` LIMIT 1
	if err != nil {
		fmt.Println("get first record not find : Error")
	}
	fmt.Println("judul", data.Judul)
	fmt.Printf("objek data %v", data)
	// ambil semua data, dlm bntk array (di go disbt slash)
	var dataset []transition.Penyimpanan
	err = db.Debug().Find(&dataset).Error //SELECT * FROM `penyimpanans`
	if err != nil {
		fmt.Println("data tidak ditemukan, Error")
	}
	// print satu persatu
	for _, d := range dataset {
		fmt.Println("judul", d.Judul)
		fmt.Printf("objek data %v", d)
	}

	// find() tdk ada limit sedangkan first() ada LIMIT 1, first() dan Take() bedanya Take() tdk ORDER & find() tdk ada ORDER
	// kita akan cari berdasarkan judul
	err = db.Debug().Where("judul = ?", "sword season 2").Find(&dataset).Error //SELECT * FROM `penyimpanans` WHERE judul = 'sword season 2'
	if err != nil {
		fmt.Println("data tidak ditemukan, Error")
	}
	// print satu persatu
	for _, d := range dataset {
		fmt.Println("judul", d.Judul)
		fmt.Printf("objek data %v", d)
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
