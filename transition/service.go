package transition

type Service interface {
	FindAll() ([]Penyimpanan, error)
	FindByID(ID int) (Penyimpanan, error)
	Create(dataRequest ItemRequest) (Penyimpanan, error)
	// define func Update dlm interface Service
	Update(ID int, dataRequest ItemRequest) (Penyimpanan, error)
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

// Implement func update dari interface service
// update hrs tahu id berapa yg akan diupdate, jd perlu param id
func (s *service) Update(ID int, dataRequest ItemRequest) (Penyimpanan, error) {
	// dptkan id yg ingin diupdate
	dataId, err := s.repository.FindByID(ID)

	rating, _ := dataRequest.Rating.Int64()
	// update nilai/datanya
	dataId.Judul = dataRequest.Judul
	dataId.Rating = int(rating)
	dataId.SubTitle = dataRequest.SubTitle

	newData, err := s.repository.Update(dataId)
	return newData, err
}
