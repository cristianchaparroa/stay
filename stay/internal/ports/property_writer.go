package ports

type PropertyWriter interface {

	// GetAllPlaces retrieves all places related to user with role Host or Manager
	GetAllPlaces(uid string) []*Property

	// GetPlace retrieves a place given a user uid and place id.
	GetPlace(uid, id string) (*Property, error)
}
