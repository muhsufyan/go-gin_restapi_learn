package transition

import (
	"fmt"

	"gorm.io/gorm"
)

// buat interface (didlmnya hanya define fungsi kosong (tdk ada implementasinya))
type Respository interface {
	// define fungsi" yg terkait dg penyimpanan (tabel penyimpanans). semuanya public (huruf awal adlh kapital)
	// fungsi get semua buku, returnnya Penyimpanan(yaitu semua data) dan error
	FindAll() ([]Penyimpanan, error)
	// get data by id dg param id (dlm bntk int), returnnya Penyimpanan (berupa 1 buah record) & error
	FindByID(ID int) (Penyimpanan, error)
	// fungsi create dg parameter data dlm (objek Penyimpanan), returnnya Penyimpanan(yaitu data yg disimpan) dan error
	Create(data Penyimpanan) (Penyimpanan, error)
}

// struct repository adlh private (huruf awal ditulis dg huruf kecil)
type repository struct {
	// perlu objek db dari main, sehingga kita lempar dari main ke sini lalu tangkap repository ini
	// tangkap objek db
	db *gorm.DB
}

// kita hrs instansiasi dulu sehingga perlu keyword New, paramnya objek dr db (agar dpt mengakses db), returnnya adlh repository
func NewRepository(db *gorm.DB) *repository {
	// passing objek dbnya, sehingga kita dpt passing objek db yg diakses oleh main.go
	// ini instansiasi dr struct repository diatas
	return &repository{db}
}

// IMPLEMENTASI DARI FUNGSI DALAM INTERFACE DIATAS, return Penyimpanan (data) dan error
func (r *repository) FindAll() ([]Penyimpanan, error) {
	// variabel yg mrpkn array/slash of Penyimpanan
	var dataset []Penyimpanan
	// get all data, dulu langsung lewat db tp now hrs lwt r dulu baru ke db. db ini berasal dari struct repository yaitu db *gorm.DB jd jika diganti jd dbase maka r.db.xxx db nya dibwh ini hrs diganti jd r.dbase.xx
	err := r.db.Find(&dataset).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return dataset, err
}

// param ID, return Penyimpanan & error
func (r *repository) FindById(ID int) (Penyimpanan, error) {
	// variabel yg mrpkn array/slash of Penyimpanan
	var data Penyimpanan
	// get all data, dulu langsung lewat db tp now hrs lwt r dulu baru ke db
	err := r.db.Find(&data, ID).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return data, err
}

func (r *repository) Create(data Penyimpanan) (Penyimpanan, error) {
	// get all data, dulu langsung lewat db tp now hrs lwt r dulu baru ke db
	err := r.db.Create(&data).Error
	if err != nil {
		fmt.Println("error, data tidak ditemukan")
	}
	return data, err
}

// kita tlh menerapkan konsep repository layer/repository pattern, asal alurnya langsung akses ke db now sbb
// dr main->repository->db->mysql
// layer repository bertanggung jwb berhub dg database
