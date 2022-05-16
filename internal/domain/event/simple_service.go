package event

import "math"

type SimpleService struct {
	repo Repository
}

func NewSimpleService(r Repository) Service {
	return SimpleService{
		repo: r,
	}
}

func (s SimpleService) FindAll() (*[]Event, error) {
	return s.repo.FindAll()
}

func (s SimpleService) FindOne(id int) (*Event, error) {
	return s.repo.FindOne(id)
}

func (s SimpleService) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	return s.repo.FindByCoords(latitude, longitude, math.Abs(radius))
}
