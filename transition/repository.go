package transition

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)
	Create(data Penyimpanan) (Penyimpanan, error)
	// define func update
	Update(data Penyimpanan) (Penyimpanan, error)
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

// implement func Update pd interface diatas
func (r *repository) Update(data Penyimpanan) (Penyimpanan, error) {
	// update(simpan update) ke db mysql
	err := r.db.Save(&data).Error
	return data, err
}
