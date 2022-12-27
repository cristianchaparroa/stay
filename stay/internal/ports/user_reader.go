package ports

type UserReader interface {
	// GetUser retrieves a user given a uid
	GetUser(uid string) (*User, error)
}
