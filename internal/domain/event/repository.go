package event

import (
	"Events_Backend_v2/cmd/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Repository interface {
	FindAll() (*[]Event, error)
	FindOne(uint64) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	CreateOne(*Event) error
	UpdateOne(*Event) error
	DeleteOne(uint64) error
}

type repository struct {
	sess db.Session
}

func NewRepository(conf *config.Config) (Repository, error) {
	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: conf.DatabaseName,
		Host:     conf.DatabaseHost,
		User:     conf.DatabaseUser,
		Password: conf.DatabasePassword,
	})

	if err != nil {
		return nil, err
	}

	return &repository{sess}, nil
}

func (repo *repository) FindAll() (*[]Event, error) {
	var events []Event
	err := repo.sess.Collection("events").Find().All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func (repo *repository) FindOne(id uint64) (*Event, error) {
	var event Event
	err := repo.sess.Collection("events").Find(id).One(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (repo *repository) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	events := make([]Event, 0)
	err := repo.sess.Collection("events").
		Find("SQRT(POW(? - latitude, 2) + POW(? - longitude, 2)) < ?", latitude, longitude, radius).
		All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func (repo *repository) CreateOne(event *Event) error {
	event.ID = 0
	err := repo.sess.Collection("events").InsertReturning(event)
	return err
}

func (repo *repository) UpdateOne(event *Event) error {
	return repo.sess.Collection("events").Find(event.ID).Update(event)
}

func (repo *repository) DeleteOne(id uint64) error {
	return repo.sess.Collection("events").Find(id).Delete()
}
