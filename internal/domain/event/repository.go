package event

type Repository interface {
	FindAll() (*[]Event, error)
	FindOne(uint64) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	PostOne(*Event) error
	UpdateOne(uint64, *Event) error
	DeleteOne(uint64) error
}
