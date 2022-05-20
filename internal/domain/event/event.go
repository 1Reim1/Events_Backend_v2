package event

type Event struct {
	Id               int      `json:"id" db:"id"`
	Title            string   `json:"title" db:"title"`
	ShortDescription string   `json:"short_description" db:"short_description"`
	Description      string   `json:"description" db:"description"`
	EventDate        string   `json:"event_date" db:"event_date"`
	Latitude         float64  `json:"latitude" db:"latitude"`
	Longitude        float64  `json:"longitude" db:"longitude"`
	Images           []string `json:"images" db:"images"`
	Preview          string   `json:"preview" db:"preview"`
}

type EventWithoutID struct {
	Title            string   `db:"title"`
	ShortDescription string   `db:"short_description"`
	Description      string   `db:"description"`
	EventDate        string   `db:"event_date"`
	Latitude         float64  `db:"latitude"`
	Longitude        float64  `db:"longitude"`
	Images           []string `db:"images"`
	Preview          string   `db:"preview"`
}

func (e *Event) WithoudID() *EventWithoutID {
	return &EventWithoutID{
		e.Title,
		e.ShortDescription,
		e.Description,
		e.EventDate,
		e.Latitude,
		e.Longitude,
		e.Images,
		e.Preview,
	}
}
