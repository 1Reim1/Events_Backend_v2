package event

import (
	"Events_Backend_v2/cmd/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type PostgresqlRepository struct {
	sess db.Session
}

func NewPostgresqlRepository(conf *config.Config) (*PostgresqlRepository, error) {
	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: conf.DatabaseName,
		Host:     conf.DatabaseHost,
		User:     conf.DatabaseUser,
		Password: conf.DatabasePassword,
	})

	if err != nil {
		return nil, err
	}

	return &PostgresqlRepository{sess}, nil
}

func (repo *PostgresqlRepository) FindAll() (*[]Event, error) {
	events := make([]Event, 0)
	err := repo.sess.Collection("events").Find().All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func (repo *PostgresqlRepository) FindOne(id int) (*Event, error) {
	event := Event{}
	err := repo.sess.Collection("events").Find(id).One(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (repo *PostgresqlRepository) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	events := make([]Event, 0)
	err := repo.sess.SQL().
		SelectFrom("events").
		Where("SQRT(POW(? - latitude, 2) + POW(? - longitude, 2)) < ?", latitude, longitude, radius).
		All(&events)
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func (repo *PostgresqlRepository) PostOne(event *Event) error {
	_, err := repo.sess.Collection("events").Insert(event.WithoudID())
	return err
}

func (repo *PostgresqlRepository) UpdateOne(id int, event *Event) error {
	res := repo.sess.Collection("events").Find(id)
	err := res.One(&Event{})
	if err != nil {
		return err
	}
	return res.Update(event.WithoudID())
}

func (repo *PostgresqlRepository) DeleteOne(id int) error {
	res := repo.sess.Collection("events").Find(id)
	err := res.One(&Event{})
	if err != nil {
		return err
	}
	return res.Delete()
}
