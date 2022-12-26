package repositories

import "stay/internal/ports"

type User struct {
	UID   string `gorm:"primaryKey"`
	Names string
}

func (u *User) ToDomain() *ports.User {
	return &ports.User{
		UID:   u.UID,
		Names: u.Names,
	}
}
