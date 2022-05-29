package event

import "math"

type Service interface {
	FindAll() (*[]Event, error)
	FindOne(id uint64) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	CreateOne(*Event) error
	UpdateOne(*Event) error
	DeleteOne(uint64) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() (*[]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id uint64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	return (*s.repo).FindByCoords(latitude, longitude, math.Abs(radius))
}

func (s *service) CreateOne(event *Event) error {
	return (*s.repo).CreateOne(event)
}

func (s *service) UpdateOne(event *Event) error {
	return (*s.repo).UpdateOne(event)
}

func (s *service) DeleteOne(id uint64) error {
	return (*s.repo).DeleteOne(id)
}
