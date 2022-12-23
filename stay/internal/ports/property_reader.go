package ports

type PropertyReader interface {

	// Save will store a property associated to Host
	Save(p *Property, uid string) (*Property, error)
}
