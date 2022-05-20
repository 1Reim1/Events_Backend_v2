package event

type Service interface {
	FindAll() (*[]Event, error)
	FindOne(id int) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	PostOne([]byte) error
	UpdateOne(int, []byte) error
	DeleteOne(int) error
}
