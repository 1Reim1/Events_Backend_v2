package event

type Event struct {
	ID               uint64   `json:"id" db:"id,omitempty"`
	Title            string   `json:"title" db:"title"`
	ShortDescription string   `json:"short_description" db:"short_description"`
	Description      string   `json:"description" db:"description"`
	EventDate        string   `json:"event_date" db:"event_date"`
	Latitude         float64  `json:"latitude" db:"latitude"`
	Longitude        float64  `json:"longitude" db:"longitude"`
	Images           []string `json:"images" db:"images"`
	Preview          string   `json:"preview" db:"preview"`
}
