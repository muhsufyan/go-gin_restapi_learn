package transition

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)
	Create(data Penyimpanan) (Penyimpanan, error)
	Update(data Penyimpanan) (Penyimpanan, error)
	// delete
	Delete(data Penyimpanan) (Penyimpanan, error)
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

func (r *repository) Update(data Penyimpanan) (Penyimpanan, error) {
	err := r.db.Save(&data).Error
	return data, err
}

// implement func Delete pd interface Repository diatas
func (r *repository) Delete(data Penyimpanan) (Penyimpanan, error) {
	// hapus data yg diinginkan di db mysql
	err := r.db.Delete(&data).Error
	return data, err
}
