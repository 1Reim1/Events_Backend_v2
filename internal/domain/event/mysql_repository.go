package event

import (
	"Events_Backend_v2/cmd/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(conf *config.Config) (*MySQLRepository, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			conf.DatabaseUser,
			conf.DatabasePassword,
			conf.DatabaseHost,
			conf.DatabaseName,
		))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLRepository{db: db}, nil
}

func (s MySQLRepository) FindAll() (*[]Event, error) {
	//Get all events
	eventRows, err := s.db.Query("SELECT * FROM `events`")
	if err != nil {
		return nil, err
	}
	//Scan events into structures
	events := s.scanEvents(eventRows)
	_ = eventRows.Close()
	//Get all images
	if len(events) > 0 {
		imageRows, err := s.db.Query("SELECT `url`, `event_id` FROM `images`")
		if err != nil {
			return nil, err
		}
		s.initializeImagesForEvents(events, imageRows)
		_ = imageRows.Close()
	}
	return &events, nil
}

func (s MySQLRepository) FindOne(id int) (*Event, error) {
	//Get event
	e := Event{}
	row := s.db.QueryRow(fmt.Sprintf("SELECT * FROM `events` WHERE `id` = %d", id))
	err := row.Scan(&e.Id, &e.Title, &e.ShortDescription, &e.Description, &e.EventDate, &e.Latitude, &e.Longitude, &e.Preview)
	if err != nil {
		return nil, err
	}
	//Get images
	imageRows, err := s.db.Query(fmt.Sprintf("SELECT `url` FROM `images` WHERE `event_id` = %d", id))
	if err != nil {
		return nil, err
	}
	images := make([]string, 0)
	for imageRows.Next() {
		var url string
		if imageRows.Scan(&url) == nil {
			images = append(images, url)
		}
	}
	_ = imageRows.Close()
	e.Images = images
	return &e, nil
}

func (s MySQLRepository) FindByCoords(latitude, longitude, radius float64) (*[]Event, error) {
	//Get events
	eventRows, err := s.db.Query(fmt.Sprintf(
		"SELECT * FROM `events` WHERE SQRT(POW(%f - `latitude`, 2) + POW(%f - `longitude`, 2)) < %f",
		latitude,
		longitude,
		radius,
	))
	if err != nil {
		return nil, err
	}
	//Scan events into structures
	events := s.scanEvents(eventRows)
	_ = eventRows.Close()
	//Get images
	if len(events) > 0 {
		imageRows, err := s.db.Query(s.buildImagesQuery(events))
		if err != nil {
			return nil, err
		}
		s.initializeImagesForEvents(events, imageRows)
		_ = imageRows.Close()
	}
	return &events, nil
}

func (s *MySQLRepository) scanEvents(eventRows *sql.Rows) []Event {
	events := make([]Event, 0)
	for eventRows.Next() {
		e := Event{}
		err := eventRows.Scan(&e.Id, &e.Title, &e.ShortDescription, &e.Description, &e.EventDate, &e.Latitude, &e.Longitude, &e.Preview)
		if err == nil {
			events = append(events, e)
		}
	}
	return events
}

func (s *MySQLRepository) buildImagesQuery(events []Event) string {
	imagesQuery := "SELECT `url`, `event_id` FROM `images` WHERE "
	for _, event := range events {
		imagesQuery += fmt.Sprintf("`event_id` = %d OR ", event.Id)
	}
	return imagesQuery[:len(imagesQuery)-4]
}

func (s *MySQLRepository) initializeImagesForEvents(events []Event, imageRows *sql.Rows) {
	//Store images by id
	imagesMap := make(map[int][]string)
	for imageRows.Next() {
		var id int
		var url string
		if imageRows.Scan(&url, &id) == nil {
			if arr, ok := imagesMap[id]; !ok {
				imagesMap[id] = []string{url}
			} else {
				imagesMap[id] = append(arr, url)
			}
		}
	}
	//Initialize events images
	for i := 0; i < len(events); i++ {
		if images, ok := imagesMap[events[i].Id]; ok {
			events[i].Images = images
		}
	}
}
