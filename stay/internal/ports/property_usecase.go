package ports

type PropertyUseCase interface {
	// GetAll retrieves all places related to user with role Host or Manager
	GetAll(uid string) ([]*Property, error)

	// GetProperty retrieves a place given a user uid and place id.
	GetProperty(uid, id string) (*Property, error)

	// Save will store a property associated to Host
	Save(p *Property) (*Property, error)
}
