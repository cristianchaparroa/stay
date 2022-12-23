package ports

// Guest represents the information associated to the person
// that will rent a property.
type Guest struct {
	// ID represent the Identification provided by each country could be a passport.
	ID     string
	Names  string
	IDType string
}
