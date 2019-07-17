package mockkkk

type DB interface {
	Hii() string
}

type Service struct {
	repo DB
}

func NewService(db DB) *Service {
	return &Service{
		repo: db,
	}
}
func (svr Service) SayHi() string {
	return svr.repo.Hii()
}
