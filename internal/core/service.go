package core

type Storage interface{}

type Service struct {
	storage Storage
}

func NewService(st Storage) *Service {
	return &Service{
		storage: st,
	}
}
