package transition

import (
	"fmt"
)

// struct ini hrs mengimplement func sprti yg ada di repository.go yaitu FindAll, FindByID, Create dan returnnya pun hrs sama
// interfacenya sama sprti interface Repository
type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (fr *fileRepository) FindAll() ([]Penyimpanan, error) {

	var dataset []Penyimpanan
	// simpan kedlm file, kita hanya pura" saja (fokusnya hanya pembuatan dan konsep interface dan implementasinya lwt struct)
	fmt.Println("find all")
	return dataset, nil
}

func (fr *fileRepository) FindByID(ID int) (Penyimpanan, error) {

	var data Penyimpanan

	fmt.Println("find by id")
	return data, nil
}

func (fr *fileRepository) Create(data Penyimpanan) (Penyimpanan, error) {
	fmt.Println("create")
	return data, nil
}

// di main.go kita tinggal buat variabel yg akan memanggil fileRepository ini
