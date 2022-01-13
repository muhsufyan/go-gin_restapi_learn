// now dr main->handler(controller)->service->repository->db->mysql
// service itu bertanggung jwb dlm bisnis logik. simple bisnis logik adlh fitur, misal fitur upload produk jd bisnis logiknya (upload produk) penjual bikin data produk
// bikin data produk lewat service lalu data tsb perlu disimpan kedlm db sehingga hrs melalui repository
// Cara buatnya hampir mirip dg repository. 1. buat interface 2. buat structnya
package transition //package book

// buat interface
type Service interface {
	// ini akan menlanjutkan tugas handler. data json di tangkap lewat handler yg ditangkap lewat input(request.go asalnya input.go ktia ganti namanya agar tdk bingung) lalu input tsb dipassing ke service ini, lalu dikirim ke entity dan akhirnya repository
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)                //(Book, error)
	Create(dataRequest ItemRequest) (Penyimpanan, error) //(bookReuest BookRequest)(Book, error)
}

// Buat struct
// bth repository sehingga akan meng-implemen interface Respository sehingga menjd polymorpism
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Penyimpanan, error) {
	dataset, err := s.repository.FindAll()
	return dataset, err
}

func (s *service) FindByID(ID int) (Penyimpanan, error) {
	data, err := s.repository.FindByID(ID)
	return data, err
}

func (s *service) Create(dataRequest ItemRequest) (Penyimpanan, error) {
	rating, _ := dataRequest.Rating.Int64()
	// mapping
	data := Penyimpanan{
		Judul:    dataRequest.Judul,
		Rating:   int(rating),
		SubTitle: dataRequest.SubTitle,
	}

	newData, err := s.repository.Create(data)
	return newData, err
}
