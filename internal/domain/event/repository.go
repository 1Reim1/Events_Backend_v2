package event

type Repository interface {
	FindAll() (*[]Event, error)
	FindOne(int) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	PostOne(*Event) error
	UpdateOne(int, *Event) error
	DeleteOne(int) error
}
