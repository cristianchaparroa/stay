package ports

const (
	Owner   UserRole = "host"    // Owner is the keeper of a Property
	Manager UserRole = "manager" // Manager is in charge to manage a property
	Viewer  UserRole = "manager" // Viewer can see specific information allowed by an Owner or Manager
)

// UserRole determines the actions that could perform the current
// user in the system.
type UserRole string

// User represent a user in the system
type User struct {
	UID  string
	Name string
	Role UserRole
}
