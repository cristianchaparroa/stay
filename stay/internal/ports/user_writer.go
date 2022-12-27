package ports

type UserWriter interface {
	// Save store a user in the system
	Save(u *User) (*User, error)
}
