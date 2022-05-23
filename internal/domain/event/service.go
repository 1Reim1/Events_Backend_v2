package event

type Service interface {
	FindAll() (*[]Event, error)
	FindOne(id uint64) (*Event, error)
	FindByCoords(float64, float64, float64) (*[]Event, error)
	PostOne([]byte) error
	UpdateOne(uint64, []byte) error
	DeleteOne(uint64) error
}
