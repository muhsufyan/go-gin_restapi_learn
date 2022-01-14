package transition

import (
	"fmt"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (fr *fileRepository) FindAll() ([]Penyimpanan, error) {

	var dataset []Penyimpanan

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
