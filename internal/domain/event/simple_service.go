package event

import (
	"encoding/json"
	"math"
)

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

func (s SimpleService) FindOne(id uint64) (*Event, error) {
	return s.repo.FindOne(id)
}

func (s SimpleService) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	return s.repo.FindByCoords(latitude, longitude, math.Abs(radius))
}

func (s SimpleService) PostOne(content []byte) error {
	event := Event{}
	err := json.Unmarshal(content, &event)
	if err != nil {
		return err
	}
	return s.repo.PostOne(&event)
}

func (s SimpleService) UpdateOne(id uint64, content []byte) error {
	event := Event{}
	err := json.Unmarshal(content, &event)
	if err != nil {
		return err
	}
	return s.repo.UpdateOne(id, &event)
}

func (s SimpleService) DeleteOne(id uint64) error {
	return s.repo.DeleteOne(id)
}
