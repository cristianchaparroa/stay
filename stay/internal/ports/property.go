package ports

// Property represents the space rent on platforms.
type Property struct {
	ID          string
	UserUID     string
	Name        string
	Description string
	Address     string
	Typ         string
	Rooms       int
}
