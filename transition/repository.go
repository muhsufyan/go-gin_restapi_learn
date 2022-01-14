package transition

import (
	"fmt"

	"gorm.io/gorm"
)

/*
pd dsrnya interface adlh kontrak, disini kita buat interface Repository yg memiliki (3) method
suatu struct yg akan mengimplement/memenuhi kontrak hrs memenuhi semua func nya, contoh struct repository akan mengimplement interface Repository jd struct hrs memiliki&memenuhi 3 buah func dari interface Repository
we use interface Repository agar dpt gunta ganti implementasi (functionnya).
implementasi dr func interface (yaitu FindAll, FindByID, Create)
func NewService(service.go) paramnya adlh interface Repository sehingga yg dpt dipassing ke service lewat NeWService adlh objek/struct apa saja selama memenuhi interface Repository
contoh disini adlh implement repository dg objek db gorm (sehingga data dpt masuk ke mysql) sehingga jika nanti we want save as file text
sehingga kita buat repository untuk file text, yg akan diimplement oleh struct. kita buat saja filenya repositoryFile.go
*/
type Repository interface {
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)
	Create(data Penyimpanan) (Penyimpanan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Penyimpanan, error) {

	var dataset []Penyimpanan

	err := r.db.Find(&dataset).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return dataset, err
}

func (r *repository) FindByID(ID int) (Penyimpanan, error) {

	var data Penyimpanan

	err := r.db.Find(&data, ID).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return data, err
}

func (r *repository) Create(data Penyimpanan) (Penyimpanan, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return data, err
}
