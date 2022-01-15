package transition

type Service interface {
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)
	Create(dataRequest ItemRequest) (Penyimpanan, error)
	Update(ID int, dataRequest ItemRequest) (Penyimpanan, error)
	// define func Delete dlm interface Service, param ID untuk get ID yg ingin dihapus
	Delete(ID int) (Penyimpanan, error)
}

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
	data := Penyimpanan{
		Judul:    dataRequest.Judul,
		Rating:   int(rating),
		SubTitle: dataRequest.SubTitle,
	}

	newData, err := s.repository.Create(data)
	return newData, err
}

func (s *service) Update(ID int, dataRequest ItemRequest) (Penyimpanan, error) {
	dataId, err := s.repository.FindByID(ID)

	rating, _ := dataRequest.Rating.Int64()

	dataId.Judul = dataRequest.Judul
	dataId.Rating = int(rating)
	dataId.SubTitle = dataRequest.SubTitle

	newData, err := s.repository.Update(dataId)
	return newData, err
}

// Implement func Delete dari interface Service
// hrs tahu id berapa yg akan dihps, jd perlu param id
func (s *service) Delete(ID int) (Penyimpanan, error) {
	// dptkan id yg ingin dihapus
	dataId, err := s.repository.FindByID(ID)
	// lempar data id yg ingin dihapus ke Repository nantinya di Repository id dr data yg ingin dihapus akan dihapus di db
	deleteData, err := s.repository.Delete(dataId)
	return deleteData, err
}
