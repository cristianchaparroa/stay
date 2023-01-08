package ports

type PropertyWriter interface {
	// Save will store a property associated to Host
	Save(p *Property) (*Property, error)

	// Delete will remove a property
	Delete(id string) error
}
